package conditionator

type ConnectorType string

const (
	And ConnectorType = "AND"
	Or  ConnectorType = "OR"
)

type Condition struct {
	Key       string
	Value     interface{}
	Connector ConnectorType
}

type Conditionator struct {
	Conditions []Condition
	query      string
}

func (c *Conditionator) Init(baseQuery string) *Conditionator {
	c.query = baseQuery
	return c
}

func (c *Conditionator) Where() *Conditionator {
	c.query += " WHERE "
	return c
}

func (c *Conditionator) Build() (string, []interface{}) {
	if c.Conditions == nil || len(c.Conditions) == 0 {
		return c.query, []interface{}{}
	}
	var params []interface{}
	for _, condition := range c.Conditions {
		connector := ""
		if len(params) > 0 {
			connector = " " + string(condition.Connector) + " "
		}
		c.query += connector + condition.Key + " = ?"
		params = append(params, condition.Value)
	}
	return c.query, params
}

func (c *Conditionator) AddCondition(condition Condition) *Conditionator {
	switch condition.Value.(type) {
	case nil:
		return c
	case string:
		if condition.Value == "" {
			return c
		}
	case int:
		if condition.Value == 0 {
			return c
		}
	}
	c.Conditions = append(c.Conditions, condition)
	return c
}
