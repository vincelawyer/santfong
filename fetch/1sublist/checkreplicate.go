package main

func checkReplicate(list productList) (rep []string) {
	m := make(map[string]bool)

	for _, item := range list.Items {
		_, ok := m[item.Title]
		if ok {
			rep = append(rep, item.Title)
		} else {
			m[item.Title] = true
		}
	}

	return
}
