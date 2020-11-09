# MongoDB

## Установка

### Качаем архив с официального сайта

https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu-tarball/

### Распаковываем архив

$ tar -zxvf mongodb-linux-*-4.4.1.tgz
$ sudo cp <mongodb-install-directory>/bin/* /usr/local/bin/

### Настройка директории с базой данных

$ sudo mkdir -p /var/lib/mongo
$ sudo chown `whoami` /var/lib/mongo

### Настройка директории с логами

$ sudo mkdir -p /var/log/mongodb
$ sudo chown `whoami` /var/log/mongodb

### Запуск серввера

$ mongod --dbpath /var/lib/mongo --logpath /var/log/mongodb/mongod.log --fork

### Запуск оболочки

$ mongo


## Работа

### Создание и начало работы с db

> use tutorial
switched to db tutorial

### Вставка записи

> db.users.insert({"name":"Petya"})
WriteResult({ "nInserted" : 1 })

### Получение записей

> db.users.find();
{ "_id" : ObjectId("5fa944f360abed1bbd995c8d"), "name" : "Petya" }

#### Запрос полей

> db.users.find({city:"Moscow"}, {name:1});

### Количество записей в коллекции 

> db.users.count();
3

### Обновление записей $set

> db.users.update({"name": "Petya"}, {
	$set:{
		country:"Russia"
	}
});

> db.users.update({"name": "Petya"}, {
	$set:{
		favourites:{
			movies:["Casablanca"]
		}
	}
});

### Удаление полей $unset

> db.users.update({"city":"Moscow"}, {
	$unset:{
		city:1
	}
})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })

### Добавление значений в списов $addToSet/$push

> db.users.update({name:"A"}, {
	$addToSet:{
		"favourites.movies":"Rambo"
	}
})

