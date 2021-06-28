package api

func (o Permission) String() string {
	ret := o.GetAction() + ":"
	r := o.GetResource()

	if r.GetOrgID() != "" {
		ret += "orgs/" + r.GetOrgID()
	}
	ret += "/" + r.GetType()
	if r.GetId() != "" {
		ret += "/" + r.GetId()
	}
	return ret
}
