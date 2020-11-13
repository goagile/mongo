# MongoDB

## Установка

### Качаем архив с официального сайта

```

	https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu-tarball/

```

### Распаковываем архив

```js
	
	tar -zxvf mongodb-linux-*-4.4.1.tgz

```

### Настройка директории с базой данных

```js

	sudo mkdir -p /var/lib/mongo
	sudo chown `whoami` /var/lib/mongo

```

### Настройка директории с логами

```js

	sudo mkdir -p /var/log/mongodb
	sudo chown `whoami` /var/log/mongodb

```

### Запуск серввера

```js

	mongod --dbpath /var/lib/mongo --logpath /var/log/mongodb/mongod.log

```

### Запуск оболочки

```

	./bin/mongo

```


## Работа

### Создание и начало работы с db

```js

	> use tutorial
	switched to db tutorial

```

### Вставка записи

```js
	
	> db.users.insert(
		{
			"name": "Ivan"
		}
	)

	WriteResult({ "nInserted" : 1 })

```

### Получение записей

```js
	
	> db.users.find();
	{ "_id" : ObjectId("5fa944f360abed1bbd995c8d"), "name" : "Ivan" }

```

#### Запрос полей

```js
	
	> db.users.find(
		{
			city: "Moscow"
		}, 
		{
			name: 1
		}
	);

```

### Количество записей в коллекции 

```js

	> db.users.count();
	3

```

### Обновление записей $set

```js

	> db.users.update(
		{
			"name": "Ivan"
		}, 
		{
			$set:{
				country:"Russia"
			}
		}
	);

```

```js

	> db.users.update(
		{
			"name": "Ivan"
		}, 
		{
			$set: {
				favourites: {
					movies: ["Casablanca"]
				}
			}
		}
	);

```

### Удаление полей $unset

```js

	> db.users.update(
		{
			"city": "Moscow"
		}, 
		{
			$unset: {
				city: 1
			}
		}
	)
	WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })

```

### Добавление значений в список $addToSet/$push

```js

	> db.users.update(
		{
			name: "A"
		}, 
		{
			$addToSet: {
				"favourites.movies": "Rambo"
			}
		}
	)

	> db.users.update(
		{
			name: "A"
		}, 
		{
			$push: {
				"favourites.movies":"Rambo"
			}
		}
	)

```

### Удаление записей

```js

	> db.users.remove(
		{
			name: "A"
		}
	)

```

### Удаление коллекций 

```js

	> db.users.drop()

```

### Добавление записей

```js

	> 	for (var i = 0; i < 200000; i++) {
	...		db.numbers.save({num:i})
	... }

```

### Поиск в диапазоне значений

```js
	
	> db.numbers.find({$gt:199995, $lt:199999})

```

### Простой Индекс

#### Профиль запроса

```js

	> db.numbers.find(
		{
			num: {
				$gt: 199995
			}
		}
	).explain()

```

#### Создание индекса 

```js

	> db.numbers.ensureIndex({num:1})

```

#### Получение индексов над коллекцией

```js

	> db.numbers.getIndexes()

```

### Информация об объектах

#### Список баз данных и коллекций

```js

	> show dbs
	> show collections

```

#### Детальная информация по базе данных или коллекции

```js

	> db.stats()
	> db.users.stats()

```

### Команды "под капотом"

#### Информация по базе данных

```js

	> db.stats()
	> db.runCommand({dbstats:1})
	> db.$cmd.findOne({dbstats:1})

```

#### Информация по коллекции

```js

	> db.users.stats()
	> db.runCommand({collstats:"users"})
	> db.$cmd.findOne({collstats:"users"})

```
