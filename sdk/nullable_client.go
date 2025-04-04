package sdk

type NullableClient struct {
	nullableResourcexClient
	nullableStaticWebsiteClient
}

var _ IClient = (*NullableClient)(nil)
