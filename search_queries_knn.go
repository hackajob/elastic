package elastic

import "fmt"

// KNNQuery allows to define KNN as filters.
type KNNQuery struct {
	query Query
	field string
	value []float32
	k     int
}

// NewKNNQuery creates and initializes a new KNNQuery.
func NewKNNQuery(
	query Query,
	field string,
	value []float32,
	k int,
) *KNNQuery {
	return &KNNQuery{
		query: query,
		field: field,
		value: value,
		k:     k,
	}
}

func (q *KNNQuery) Source() (interface{}, error) {
	source := make(map[string]map[string]interface{})
	field := make(map[string]interface{})

	src, err := q.query.Source()
	if err != nil {
		return nil, err
	}
	field["vector"] = q.value
	field["k"] = q.k
	field["filter"] = src

	source["knn"] = map[string]interface{}{}
	source["knn"][q.field] = field

	fmt.Println(fmt.Sprintf("%+v", source))
	return source, nil
}
