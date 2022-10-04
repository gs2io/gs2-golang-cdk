package core

import "strings"

type GetAttr struct {
	key string
}

func NewGetAttrByResource(
	resource ICdkResource,
	path string,
) GetAttr {
	return GetAttr{
		key: resource.ResourceName() + "." + path,
	}
}

func NewGetAttrByKey(
	key string,
) GetAttr {
	return GetAttr{
		key: key,
	}
}

func (p GetAttr) String() string {
	return "!GetAttr " + p.key
}

func NewGetAttrRegion() GetAttr {
	return GetAttr{
		key: "Gs2::Region",
	}
}

func NewGetAttrOwnerId() GetAttr {
	return GetAttr{
		key: "Gs2::OwnerId",
	}
}

type Join struct {
	separator string
	values    []string
}

func NewJoin(
	separator string,
	values []string,
) Join {
	return Join{
		separator: separator,
		values:    values,
	}
}

func (p Join) String() string {
	return "!Join ['" + p.separator + "', [" + strings.Join(p.values, ",") + "]]"
}
