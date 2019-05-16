func main() {
	app := "top"
	arg0 := "-b"
	arg1 := "-n"
	arg2 := "1"

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
	  println(err.Error())
	  return
	}

	print(string(stdout))
}