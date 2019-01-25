package models

import (
	"errors"
)

// go:generate easyjson -all $GOFILE
type Account struct {
	// адрес электронной почты пользователя. Тип - unicode-строка длиной до 100 символов.
	// Гарантируется уникальность
	Email string `json:"email"`

	// имя и фамилия соответственно. Тип - unicode-строки длиной до 50 символов.
	// Поля опциональны и могут отсутствовать в конкретной записи.
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`

	// номер мобильного телефона. Тип - unicode-строка длиной до 16 символов.
	// Поле является опциональным, но для указанных значений гарантируется уникальность. Заполняется довольно редко.
	Phone string `json:"phone"`

	// unicode-строка "m" означает мужской пол, а "f" - женский.
	Gender string `json:"sex"`

	// дата рождения, записанная как число секунд от начала UNIX-эпохи по UTC
	// (другими словами - это timestamp). Ограничено снизу 01.01.1930 и сверху
	// 01.01.1999-ым.
	BirthDate int64 `json:"birth"`

	// страна проживания. Тип - unicode-строка длиной до 50 символов. Поле опционально.
	Country string `json:"country"`

	// город проживания. Тип - unicode-строка длиной до 50 символов.
	// Поле опционально и указывается редко. Каждый город расположен в определённой стране.
	City string `json:"city"`

	//  дата регистрации в системе. Тип - timestamp с ограничениями: снизу 01.01.2011, сверху 01.01.2018.
	Joined int64 `json:"joined"`

	// текущий статус пользователя в системе. Тип - одна строка из следующих вариантов:
	// "свободны", "заняты", "всё сложно". Не обращайте внимание на странные окончания :)
	Status string `json:"status"`

	// интересы пользователя в обычной жизни. Тип - массив unicode-строк, возможно пустой.
	// Строки не превышают по длине 100 символов.
	Interests []string `json:"interests"`

	// начало и конец премиального периода в системе (когда пользователям очень хотелось найти "вторую половинку"
	// и они делали денежный вклад). В json это поле представлено вложенным объектом с полями start и finish,
	// где записаны timestamp-ы с нижней границей 01.01.2018.
	Premium Premium `json:"premium"`

	// массив известных симпатий пользователя, возможно пустой.
	// Все симпатии идут вразнобой и каждая представляет собой объект из следующих полей:
	//id - идентификатор другого аккаунта, к которому симпатия. Аккаунт по id в исходных данных всегда существует.
	//     В данных может быть несколько лайков с одним и тем же id.
	//ts - время, то есть timestamp, когда симпатия была записана в систему.
	Likes []Like `json:"likes"`

	// уникальный внешний идентификатор пользователя. Устанавливается
	// тестирующей системой и используется затем, для проверки ответов сервера.
	// 32-разрядное целое число.
	Id uint32 `json:"id"`
}

type Premium struct {
	Start int64 `json:"start"`
	Finish int64 `json:"finish"`
}

type Like struct {
	Id int32 `json:"id"`
	Time int64 `json:"ts"`
}


func (v *Account) GetId() uint32 {
	return v.Id
}

func (v *Account) Validate() error {
	switch {
	case v.Id == 0:
		return errors.New("id should be non-zero")
	case len(v.Email) > 100:
		return errors.New("email is too long")
	case len(v.FirstName) > 50:
		return errors.New("first_name is too long")
	case len(v.LastName) > 50:
		return errors.New("last_name is too long")
	case len(v.Phone) > 16:
		return errors.New("phone is too long")
	case v.Gender != "m" && v.Gender != "f":
		return errors.New("invalid gender")
	case v.BirthDate < -1262304000 || v.BirthDate > 1514851199:
		//log.Printf("invalid birth_date: %d", v.BirthDate)
		return errors.New("invalid birth_date")
	case len(v.Country) > 50:
		return errors.New("country is too long")
	case len(v.City) > 50:
		return errors.New("country is too long")
	case v.Joined < 1293840000 || v.Joined > 1514851199:
		return errors.New("joined date is invalid")
	case v.Status != "свободны" && v.Status != "заняты" && v.Status != "всё сложно":
		return errors.New("status is invalid")
	case v.Premium.Start < 1514764800 && v.Premium.Start > v.Premium.Finish:
		return errors.New("premium is invalid")
	}

	for _, v := range v.Interests {
		if len(v) > 100 {
			return errors.New("interest: \"" + v + "\" is too long")
		}
	}
	return nil
}

func (v Account) IsValid() bool {
	return v.Id != 0
}
