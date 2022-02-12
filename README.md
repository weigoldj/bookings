# Bookings and Reservations

A sample bookings and reservations website used to learn about building web 
sites in Go lang from the course "Building Modern Web Applications with Go (Golang)"


Tech Stack
- Built in Go version 1.17
- [chi router](github.com/go-chi/chi/v5) - lightweight, idiomatic and composable router
- [Alex Edwars SCS](github.com/alexedwards/scs/v2) - HTTP Session Management for Go
- [nosurf](github.com/justinas/nosurf) - Cross-Site Request Forgery attacks
- [bootstrap](https://getbootstrap.com) - css / javascript 
- [bootstrap date picker](https://bootstrap-datepicker.readthedocs.io/en/latest/#) - calendar to pick dates
- [notie](https://github.com/jaredreich/notie) - top alert bar 
- [SweetAlert2](https://sweetalert2.github.io) - modal alerts
Database
- [Soda / pop](https://gobuffalo.io/en/docs/db/getting-started/)
- [fizz](https://gobuffalo.io/en/docs/db/fizz/)


### Soda migrations
Soda is part of buffalo but you can download and set path to use it.  Soda uses Fizz a DSL to migrate databases.  

```
# create fizz migration
> soda generate fizz <migration name>
'''
create_table("users") {
  t.Column("id", "integer", {primary: true})
  t.Column("first_name", "string", { "default": "" })
  t.Column("last_name", "string", { "default": "" })
  t.Column("email", "string", {})
  t.Column("password", "string", {"size": 60})
  t.Column("access_level", "integer", {"default": 1})
}
'''
# create indexes
soda generate fizz CreateUniqueInexOnTable 
'''
add_index("users", "email", {"unique": true })

drop_index("users", "users_email_idx")
'''


### Generate Index from one tabe to another 
```
add_foreign_key( "from_table", "from_table_id", {"to_table": ["id"]})
    "on_delete": "cascade"
    "on_update": "cascade"
})

drop_)foreignt_key("from_table", "from_table_to_table_id_fx, {} )
```
# run migration
> soda migrate 
> soda migrate down
> soda reset  # will run all all down migrations, then run all migrations.
# help
> soda help
```