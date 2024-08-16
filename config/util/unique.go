package util

// Unique 数组去重
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{}) // 使用空结构体作为值，节省内存
	var result []T

	for _, value := range slice {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
