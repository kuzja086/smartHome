# smartHome
Backend для системы умного дома
# Цели
* Разработать собственную систему умного дома, которая будет использоваться в реальных условиях;
* Изучить язык [Golang](https://go.dev/);
* Изучить смежные технологии, методики, попробовать что-то новое;
* Разработать Web-интерфейс для бекэнда;
* Подулючить к своему Умному дому различные умные устройства с использованием raspbery и arduino;
# Технологии
В некототорых местах проекта, безусловно присутсвует ***"Overengineering"***, но т.к. одна из целей проекта это изучение чего-то нового, то это допускается.
* Код в проекте старался писать с использованием **["Чистой архитектуры"](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)** TODO(ссылка). Отделил транспортный слой, от слоя бизнесс-логики и от слоя работы с БД;
* На текущем этапе общаться с сервером можно через **REST-интерфейс**, Документация к которому описывается спецификации **OpenApi 3.0**. Для тестирования можно использовать **Postman**, есть коллекция для использования;
* Для запуска сервера используется **Docker** и **Docker-compose**;
* На каждый push, workflow с использованием **GitHub Action**. Раннер GHA поднят на собственном серере из старого ноутбука с linux mint.
В GHA запускается сборка приложения, тестирование, сборка Docker-образа и публикация его на **Docker-hub**;
## Используемые билотеки Go
* [cleanEnv](http://github.com/ilyakaznacheev/cleanenv)
* [httprouter](http://github.com/julienschmidt/httprouter)
# Установка
# Использование
# Лицензия
    Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
# Дополнительно
Пожелания и ошибки можно регистрировать в Issues.