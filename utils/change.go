package utils

import "github.com/aws/aws-sdk-go/service/dynamodb"

func Change(t *dynamodb.AttributeValue) string {
	// handle t type
	result := ""
	if t.B != nil {
		return string(t.B)
	} else if t.BS != nil {
		for _, b := range t.BS {
			result += string(b)
		}
		return result
	} else if t.L != nil {
		for _, v := range t.L {
			result += Change(v) + "_"
		}
		return result
	} else if t.BOOL != nil {
		if *t.BOOL {
			return "true"
		} else {
			return "false"
		}
	} else if t.M != nil {
		// "map[string]*AttributeValue"
		for k, v := range t.M {
			result += k + Change(v)
		}
		return result
	} else if t.N != nil {
		return *t.N
	} else if t.NS != nil {
		for _, s := range t.NS {
			result += *s + "_"
		}
		return result
	} else if t.SS != nil {
		for _, s := range t.NS {
			result += *s + "_"
		}
		return result
	} else if t.S != nil {
		return *t.S
	}
	return ""
}

func UnChange(v, t string) *dynamodb.AttributeValue {
	if t == "S" {
		return &dynamodb.AttributeValue{
			S: &v,
		}
	}
	if t == "N" {
		return &dynamodb.AttributeValue{
			N: &v,
		}
	}
	// handle t type
	//result := ""
	//if t.B != nil {
	//	return string(t.B)
	//} else if t.BS != nil {
	//	result := ""
	//	for _, b := range t.BS {
	//		result += string(b)
	//	}
	//	return result
	//} else if t.L != nil {
	//	for _, v := range t.L {
	//		result += Change(v) + "_"
	//	}
	//	return result
	//} else if t.BOOL != nil {
	//	if *t.BOOL {
	//		return "true"
	//	} else {
	//		return "false"
	//	}
	//} else if t.M != nil {
	//	// "map[string]*AttributeValue"
	//	for k, v := range t.M {
	//		result += k + Change(v)
	//	}
	//	return result
	//} else if t.N != nil {
	//	return *t.N
	//} else if t.NS != nil {
	//	for _, s := range t.NS {
	//		result += *s + "_"
	//	}
	//	return result
	//} else if t.SS != nil {
	//	for _, s := range t.NS {
	//		result += *s + "_"
	//	}
	//	return result
	//} else if t.S != nil {
	//	return *t.S
	//}
	k := ""
	return &dynamodb.AttributeValue{S: &k}
}
