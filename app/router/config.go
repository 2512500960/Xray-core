package router

import (
	"regexp"
	"strings"

	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/uuid"
	"github.com/xtls/xray-core/features/outbound"
	"github.com/xtls/xray-core/features/routing"
)

type Rule struct {
	Tag           string
	TargetTag     string
	DomainMatcher string
	Balancer      *Balancer
	Condition     Condition
}

func (r *Rule) GetTargetTag() (string, error) {
	if r.Balancer != nil {
		return r.Balancer.PickOutbound()
	}
	return r.TargetTag, nil
}

// Apply checks rule matching of current routing context.
func (r *Rule) Apply(ctx routing.Context) bool {
	return r.Condition.Apply(ctx)
}

// RestoreRoutingRule Restore implements Condition.
func (r *Rule) RestoreRoutingRule() interface{} {
	rule := r.Condition.RestoreRoutingRule().(*RoutingRule)
	rule.Tag = r.Tag
	rule.DomainMatcher = r.DomainMatcher
	if r.Balancer != nil {
		rule.TargetTag = &RoutingRule_BalancingTag{
			BalancingTag: rule.Tag,
		}
	} else {
		rule.TargetTag = &RoutingRule_OutboundTag{
			OutboundTag: rule.Tag,
		}
	}

	return rule
}

func (rr *RoutingRule) BuildCondition() (Condition, error) {
	conds := NewConditionChan()

	if len(rr.Domain) > 0 {
		switch rr.DomainMatcher {
		case "linear":
			matcher, err := NewDomainMatcher(rr.Domain)
			if err != nil {
				return nil, newError("failed to build domain condition").Base(err)
			}
			conds.Add(matcher)
		case "mph", "hybrid":
			fallthrough
		default:
			matcher, err := NewMphMatcherGroup(rr.Domain)
			if err != nil {
				return nil, newError("failed to build domain condition with MphDomainMatcher").Base(err)
			}
			newError("MphDomainMatcher is enabled for ", len(rr.Domain), " domain rule(s)").AtDebug().WriteToLog()
			conds.Add(matcher)
		}
	}

	if len(rr.UserEmail) > 0 {
		conds.Add(NewUserMatcher(rr.UserEmail))
	}

	if len(rr.InboundTag) > 0 {
		conds.Add(NewInboundTagMatcher(rr.InboundTag))
	}

	if rr.PortList != nil {
		conds.Add(NewPortMatcher(rr.PortList, false))
	} else if rr.PortRange != nil {
		conds.Add(NewPortMatcher(&net.PortList{Range: []*net.PortRange{rr.PortRange}}, false))
	}

	if rr.SourcePortList != nil {
		conds.Add(NewPortMatcher(rr.SourcePortList, true))
	}

	if len(rr.Networks) > 0 {
		conds.Add(NewNetworkMatcher(rr.Networks))
	} else if rr.NetworkList != nil {
		conds.Add(NewNetworkMatcher(rr.NetworkList.Network))
	}

	if len(rr.Geoip) > 0 {
		cond, err := NewMultiGeoIPMatcher(rr.Geoip, false)
		if err != nil {
			return nil, err
		}
		conds.Add(cond)
	} else if len(rr.Cidr) > 0 {
		cond, err := NewMultiGeoIPMatcher([]*GeoIP{{Cidr: rr.Cidr}}, false)
		if err != nil {
			return nil, err
		}
		conds.Add(cond)
	}

	if len(rr.SourceGeoip) > 0 {
		cond, err := NewMultiGeoIPMatcher(rr.SourceGeoip, true)
		if err != nil {
			return nil, err
		}
		conds.Add(cond)
	} else if len(rr.SourceCidr) > 0 {
		cond, err := NewMultiGeoIPMatcher([]*GeoIP{{Cidr: rr.SourceCidr}}, true)
		if err != nil {
			return nil, err
		}
		conds.Add(cond)
	}

	if len(rr.Protocol) > 0 {
		conds.Add(NewProtocolMatcher(rr.Protocol))
	}

	if len(rr.Attributes) > 0 {
		configuredKeys := make(map[string]*regexp.Regexp)

		for key, value := range rr.Attributes {
			configuredKeys[strings.ToLower(key)] = regexp.MustCompile(value)
		}

		conds.Add(&AttributeMatcher{configuredKeys})
	}

	if conds.Len() == 0 {
		return nil, newError("this rule has no effective fields").AtWarning()
	}

	return conds, nil
}

// Build RoutingRule translates into Rule.
func (rr *RoutingRule) Build(r *Router) (*Rule, error) {
	tag := rr.Tag
	if len(tag) == 0 {
		u := uuid.New()
		tag = u.String()
	}
	rule := &Rule{
		Tag:           tag,
		DomainMatcher: rr.DomainMatcher,
	}

	btag := rr.GetBalancingTag()
	if len(btag) > 0 {
		brule, found := r.balancers[btag]
		if !found {
			return nil, newError("balancer ", btag, " not found")
		}
		rule.Balancer = brule
	} else {
		rule.TargetTag = rr.GetTargetTag().(*RoutingRule_OutboundTag).OutboundTag
	}

	cond, err := rr.BuildCondition()
	if err != nil {
		return nil, err
	}

	rule.Condition = cond
	return rule, nil
}

func (br *BalancingRule) Build(ohm outbound.Manager) (*Balancer, error) {
	switch br.Strategy {
	case "leastPing":
		return &Balancer{
			selectors: br.OutboundSelector,
			strategy:  &LeastPingStrategy{},
			ohm:       ohm,
		}, nil
	case "random":
		fallthrough
	default:
		return &Balancer{
			selectors: br.OutboundSelector,
			strategy:  &RandomStrategy{},
			ohm:       ohm,
		}, nil

	}
}
