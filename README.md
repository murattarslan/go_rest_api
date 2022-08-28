# Go RestApi

Go ile rest api programlamak [fiber](https://github.com/gofiber/fiber) kütüphanesi ile aşırı kolay ve kapsamlı.

## içerik 

- [Test ortamı](https://github.com/murattarslan/go_router/master?readme=1#test-ortam%C4%B1-kurma)
- Router kurulumu ve dinleme(https://github.com/murattarslan/go_router/master?readme=1#router-kurulumu-ve-dinleme)
- RestApi fonksiyon tipleri
  - Get
  - Post
  - Put
  - Del
- İstekten veri alma yöntemleri
  - Params
  - Body
  - Query
- Sık kullanılan status kodları
  - 200 OK
  - 201 Created
  - 400 Bad Request
  - 401 Unauthorized
  - 403 Forbidden
  - 404 Not Found
  - 500 Internal Server Error
  - 503 Service Unavailable
  
### Test ortamı kurma

1. Localhost yardımıyla isteklerin atılıp dinlenebileceği bir sunucu kurulmalı. Bu işlem için [XAMPP](https://www.apachefriends.org/tr/) yazılımından destek alınabilir.
 
 Tabi xampp kullanmadan yapma yöntemleri de mevcut. Mesela mac işletim sisteminde ``` sudo apachectl start ``` kodunu terminale yazıp çalıştırmak yeterli olacaktır.
 
 2. Sunucuya istekleri kolayca gönderip sonuçları görebilmek için ise [Postman](https://www.postman.com/) gerekli tüm yardımı sağlayacaktır. Hem ücretsiz hem de pratik.
 
 
 ### Router kurulumu ve dinleme
 
 Pratikliği ve kapsamlı fonksiyonları dolayısıyla tercihim [fiber](https://github.com/gofiber/fiber).
 
 import ederek başlayalım.
 ```
 import (
	"github.com/gofiber/fiber/v2"
)
 ```
 
 kurulumu genel olarak şu şekilde;
 ```
 app := fiber.New()

	app.Get("/get-all-base", getAllBase)

	app.Listen(":8080")
 ```
 
 halen bir hata durumunda bilgi alabilmek için dinleme fonksiyonunu şu şekilde çağırmamız öneriliyor. ``` log.Fatal(app.Listen(":8080")) ```
 
 ### RestApi fonksiyon tipleri
 
 Rest servisleri dört sebeple kullanılır.
 - Veri eklenecektir
 - Veri düzenlenecektir
 - Veri silinecektir
 - Veriler görüntülenecektir
 
 Her biri için birer yöntem mevcut.
 
 #### Get
 
 İstemciler veri görüntülemek istediğinde verileri sunucudan almak ister. Bu işlem için 'Get' metodu mevcuttur.
 
 ```
 app.Get("/get-all-base", getAllBase)
 ```
 
 #### Post
 
 İstemciler veri eklemek istediğinde sunucuya bu yeni veriyi göndermek ister. Bu işlem için 'Post' metodu mevcuttur.
 
 ```
 app.Post("/add-base", addBase)
 ```
 
 #### Put
 
 İstemciler halihazırda mevcut bir veriyi güncellemek istediğinde sunucuya bu verinin yeni halini göndermek ister. Bu işlem için 'Put' metodu mevcuttur.
 
 :warning: Bu fonksiyonda url yolunda objeyi doğru şekilde bulabilmek için bir adet id olur.
 
 ```
 app.Put("/update-base/:id", updateBase)
 ```
 
 #### Delete
 
 İstemciler halihazırda mevcut bir veriyi silmek istediğinde sunucuyu bundan haberdar etmek ister. Bu işlem için 'Delete' metodu mevcuttur.
 
 :warning: Bu fonksiyonda url yolunda objeyi doğru şekilde bulabilmek için bir adet id olur.
 
 ```
 app.Delete("/delete-base/:id", deleteBase)
 ```
 
 Fonksiyonlar parametre olarak url yolu ve handler fonksiyonu alıyor. 
 - Url yoluna temel url adresinden sonraki kısımdır.
 - Handler fonksiyonu parametre olarak fiber kütüphanesinden bir obje alıyor. return değeri olaraksa error objesi dönüyor.
 
 örnek bir handler fonksiyonu göstermek gerekirse;
 ```
func getBase(c *fiber.Ctx) error {
        .
        .
        .
	jsonString, err := json.Marshal(response)
	if err != nil {
		return c.Status(500).SendString(string(err.Error()))
	}

	return c.Status(200).SendString(string(jsonString))
}
 ```
 
### İstekten veri alma yöntemleri

 Duruma göre istemci tüm verileri istemek yerine; tüm verilerin bir kısmını, ya da belli şartları sağlayan verileri almak ya da bu verilerde işlem yapmak isteyebilir.
 Bu gibi durumlarda servisler veri de alabilir.
 
 Mesela Bir veri eklenmek isteniyor ise bu eklenecek veri 'body' ile servise gönderilir. Ya da 'params' veya 'query' ile bir id göndererek onu silmek isteyebiliriz.
 
  #### Params
  
  Bu tip veri transferi url üzerinden olur, bu sebeple 'params' üzerinden çok büyük veriler gönderilmez.
  
  örnek olarak [Put] işleminde url yolu olarak ``` "/update-base/:id" ``` gönderdik. Burada id bir 'params' değeridir.
  
  ```
	id := c.Params("id")
        .
        .
        .
	return c.Status(200).SendString(string(jsonString))
  
  ```
  
  #### Body
  
  Eğer istemci bir obje eklemek istiyorsa bu objeyi en pratik alma şekli net 'body'.
  
  ```
	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString(string(err.Error()))
	}
          .
          .
          .
	return c.Status(200).SendString(string(jsonString))
  
  ```
  
  #### Query
  
  İsteğe bağlı sorgu kısıtlamaları için 'query' kullanılabilir. Örneğin; limit, sıralama parametresi gibi.
  
  ```
	limit := c.Query("limit")
        .
        .
        .
	return c.Status(200).SendString(string(jsonString))
  
  ```
  
### Sık kullanılan status kodları
  #### 200 OK
  #### 201 Created
  #### 400 Bad Request
  #### 401 Unauthorized
  #### 403 Forbidden
  #### 404 Not Found
  #### 500 Internal Server Error
  #### 503 Service Unavailable
 
 
 
 
 
 
 
 