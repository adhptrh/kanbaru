package path

var Paths []map[string]interface{} = []map[string]interface{}{}

func All(path string, target interface{}) {
	Paths = append(Paths, map[string]interface{}{
		"path":   path,
		"target": target,
		"method": "ALL",
	})
}

func Get(path string, target interface{}) {
	Paths = append(Paths, map[string]interface{}{
		"path":   path,
		"target": target,
		"method": "GET",
	})
}

func Post(path string, target interface{}) {
	Paths = append(Paths, map[string]interface{}{
		"path":   path,
		"target": target,
		"method": "POST",
	})
}

func Put(path string, target interface{}) {
	Paths = append(Paths, map[string]interface{}{
		"path":   path,
		"target": target,
		"method": "PUT",
	})
}

func Patch(path string, target interface{}) {
	Paths = append(Paths, map[string]interface{}{
		"path":   path,
		"target": target,
		"method": "PATCH",
	})
}

func Delete(path string, target interface{}) {
	Paths = append(Paths, map[string]interface{}{
		"path":   path,
		"target": target,
		"method": "DELETE",
	})
}
