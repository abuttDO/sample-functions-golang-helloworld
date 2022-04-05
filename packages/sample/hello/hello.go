package main

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "Ajnabee"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"
	return msg
}
