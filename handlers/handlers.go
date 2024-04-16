package handler

func getUsers(c *gin.Context) {

	var users = Users{}

	setUsersFromFile(&users, "users.json")

	if err := c.ShouldBindWith(&users, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Успешно", "result": users})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error(), "result": "[]"})
	}
}

func getUser(c *gin.Context) {
}

func addUser(c *gin.Context) {
}

func deleteUser(c *gin.Context) {
}

func setUsersFromFile(u *Users, file string) {

	pwd, _ := os.Getwd()
	// Чтение содержимого файла
	fileData, err := ioutil.ReadFile(pwd + string(os.PathSeparator) + file)

	if err != nil {
		log.Fatal("ERR_FILE. Ошибка чтения файла: ", err)
	}

	err = nil

	err = json.Unmarshal([]byte(string(fileData)), &u)

	if err != nil {
		log.Fatal("ERR_JSON. ошибка распознавания json: ", err)
	}
}
