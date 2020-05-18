package gql

type Schema struct {
	Query
	Mutation
	datasets *datasetRegistry
}

func NewSchema() *Schema {
	schema := &Schema{}
	schema.datasets = NewDatasetRegistry()
	schema.Query.datasets = schema.datasets
	schema.Mutation.datasets = schema.datasets
	return schema
}

const SchemaString = `
enum PropertyValueType {
    INTEGER
    REAL
    PERCENTAGE
    VERSION
    TIMESTAMP
    TYPE
    FILE
}

schema {
    query: Query
    mutation: Mutation
}

type Query {
    object(datasetId: ID!, id: ID!): Object!
	datasets(Ids: [ID!]): [Dataset!]!
	plugins(url: String): [Plugin!]!
}

type Mutation {
	addDataset(plugin: String!, url: String!): Dataset!
	removeDataset(id: String!): Boolean!
}

type Object {
    type:                           ObjectType!
	id:								ID!
    name:                           String!
    children(typeNames: [String!]): [ObjectGroup!]!
    properties(typeNames: [String!]):  [Property!]!
}

type ObjectGroup {
    type: ObjectType!
    objects:  [Object!]!
    total:    Int!
}

type ObjectType {
    name:           String!
    pluralName: 	String!
    description:    String!
    properties:        [PropertyType!]!
    defaultProperties: [String!]!
}

type Property {
    type:  PropertyType!
    value: String!
}

type PropertyType {
    name:           String!
    valueType:      PropertyValueType!
    description:    String!
}

type Dataset {
	id:   ID!
    root: Object!
}

type Plugin {
	name: String!
	description: String!
}
`