package server

func Init(){
	Migrate()
	r := NewRouter()
	r.Run(":8000")
}