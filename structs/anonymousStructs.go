package main

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}

	// You can even nest anonymous structs as fields within other structs:

	type car struct {
		Make   string
		Model  string
		Height int
		Width  int
		// Wheel is a field containing an anonymous struct
		Wheel struct {
			Radius   int
			Material string
		}
	}
	/*
		WHEN SHOULD YOU USE AN ANONYMOUS STRUCT?
		In general, prefer named structs. Named structs make it easier to read and understand your code,
		and they have the nice side-effect of being reusable.
		I sometimes use anonymous structs when I know I won't ever need to use a struct again.
		For example, sometimes I'll use one to create the shape of some JSON data in HTTP handlers.

		If a struct is only meant to be used once,
		then it makes sense to declare it in such a way that developers down the
		road won’t be tempted to accidentally use it again.
	*/
}
