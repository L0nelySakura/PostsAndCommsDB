ЗАПУСК
1. Скачайте проект
2. Для запуска сервера запустите через CMD команды: docker-compose up --build
4. Зайдите на http://localhost:8080/ для дальнейшей работы

ДОБАВЛЕНИЕ ДАННЫХ
1. Создания поста:
mutation {
  createPost(title: "Заголовок", author: "Автор", content: "Содержание статьи", commentsEnabled: true) {
    id
    title
    author
    content
    commentsEnabled
  }
}
По умолчанию: commentsEnabled = true
2. Добавления комментария:
mutation {
  createComment(content: "Содержание комментария", author: "Комментатор", postid: "ID статьи", parentid: "ID комментария, на который ссылается текущий комментарий") {
    id
    content
    author
    postid
    parentid
  }
}
Если комментарий не ссылается ни на какой комментарий (не является ответом), то необходимо оставить поле parentid пустым.

ВЫВОД ДАННЫХ
1. Вывод всех постов:
query {
  posts {
    id
    title
    author
    content
    commentsEnabled
  }
}
Команда выведет список всех постов, хранящихся в хранилище
2. Вывод определённого поста:
query {
  post(id: "ID поста") {
    id
    title
    author
    content
    commentsEnabled
  }
}
Команда выведет пост, с указанным id
3. Вывод комментариев к определённому к посту:
query {
  comments(id: "ID поста", limit: 10, offset: 0) {
    id
    author
    content
    postid
    parentid
  }
}
Данный запрос выводит комментарии к посту с указанным id и все ответы к ним
Количество комментариев к посту определяет параметр limit, offset отвечает за количество пропускаемых комментариев
По умолчанию: limit=10, offset=0

ОПИСАНИЕ
В проекте реализована система для добавления и чтения постов с комментариями (примеры: Хабр, Reddit)
Комментарии организованы иерархически, можно добавлять ответы к комментариям
Проект написан с использованием gqlgen и PostgreSQL
