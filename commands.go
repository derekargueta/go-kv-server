package main

const (
	GET_USAGE    = "Invalid use of get:\nUsage: get [key]"
	ADD_USAGE    = "Invalid usage of add:\nUsage: add [key] [value]"
	DELETE_USAGE = "Invalid usage of delete:\nUsage: delete [key]"
)

func AddCmd(tokens []string) string {
	if len(tokens) < 3 {
		return ADD_USAGE
	}
	key := tokens[1]
	value := tokens[2]
	GetStore().Add(key, value)
	return "Added!\n"
}

func GetCmd(tokens []string) string {
	if len(tokens) != 2 {
		return GET_USAGE
	}
	key := tokens[1]
	result := GetStore().Get(key) + "\n"
	return result
}

func DeleteCmd(tokens []string) string {
	if len(tokens) != 2 {
		return DELETE_USAGE
	}
	GetStore().Delete(tokens[1])
	return "Deleted!\n"
}
