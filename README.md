# Goblog

This app serves to be a journey through how to write a highly performant, well designed, available, resilient, and secure
API that is connected to a postgres database. The intention is not to reinvent the wheel and be clever, rather it is 
write code that compiles, satisfies the requirements, and ultimately is easily readable by other humans. Error messages 
should be explicit and describe every conceivable negative path, system boundaries (such as between the database and the
app layer) should be thoroughly vetted with integration tests, and TDD best practices should be followed. In short this
app should be "good", a solid working CRUD dashboard.
