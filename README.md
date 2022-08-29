# Go RestApi

Go ile rest api programlamak [fiber](https://github.com/gofiber/fiber) kütüphanesi ile aşırı kolay ve kapsamlı. 
Bu repo tamamen eğitim amaçlı yazılmıştır. Genel hatlarıyla rest api kavramından da söz edildi. Status kodlar, rest api methodları ve handler kavramları işlendi. 

## içerik 

- [Test ortamı](https://github.com/murattarslan/go_rest_api/tree/master#test-ortam%C4%B1-kurma)
- [Router kurulumu ve dinleme](https://github.com/murattarslan/go_rest_api/tree/master#router-kurulumu-ve-dinleme)
- [RestApi fonksiyon tipleri](https://github.com/murattarslan/go_rest_api/tree/master#restapi-fonksiyon-tipleri)
  - [Get](https://github.com/murattarslan/go_rest_api/tree/master#get)
  - [Post](https://github.com/murattarslan/go_rest_api/tree/master#post)
  - [Put](https://github.com/murattarslan/go_rest_api/tree/master#put)
  - [Del](https://github.com/murattarslan/go_rest_api/tree/master#del)
- [İstekten veri alma yöntemleri](https://github.com/murattarslan/go_rest_api/tree/master#i%CC%87stekten-veri-alma-y%C3%B6ntemleri)
  - [Params](https://github.com/murattarslan/go_rest_api/tree/master#params)
  - [Body](https://github.com/murattarslan/go_rest_api/tree/master#body)
  - [Query](https://github.com/murattarslan/go_rest_api/tree/master#query)
- [Sık kullanılan status kodları](https://github.com/murattarslan/go_rest_api/tree/master#s%C4%B1k-kullan%C4%B1lan-status-kodlar%C4%B1)
  - [200 OK](https://github.com/murattarslan/go_rest_api/tree/master#200-ok)
  - [201 Created](https://github.com/murattarslan/go_rest_api/tree/master#201-created)
  - [400 Bad Request](https://github.com/murattarslan/go_rest_api/tree/master#400-bad-request)
  - [401 Unauthorized](https://github.com/murattarslan/go_rest_api/tree/master#401-unauthorized)
  - [403 Forbidden](https://github.com/murattarslan/go_rest_api/tree/master#403-forbidden)
  - [404 Not Found](https://github.com/murattarslan/go_rest_api/tree/master#404-not-found)
  - [500 Internal Server Error](https://github.com/murattarslan/go_rest_api/tree/master#500-internal-server-error)
  - [503 Service Unavailable](https://github.com/murattarslan/go_rest_api/tree/master#503-service-unavailable)
  
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

Rest api konusunda response kadar önemli bir diğer husus status code yapısıdır. 
Bu kodlar dikkate alınmadan 200 kodu ile response verebilir mi? cevap evet. Ama client sadece bu kodlara bakarak response body kontrolüne gerek bile duymadan işlem yapabilir. 

Status code [fiber](https://github.com/gofiber/fiber) kütüphanesinde iki şekilde verilebilir.
1. Status()
 ```
 return c.Status(500).SendString(string(err.Error()))
 
 ```
 Bu yöntemde response göndermeden önce status kod ekliyoruz.
 
2. SendStatus()
 ```
 c.SendStatus(500)
 return c.SendString(string(err.Error()))
 ```
 Bu yöntemde ise status kod ayrı response ayrı gönderiliyor.
 
Aşağıda en sık kullanılan bir kaç status kodu inceledik.

  ##### 200 OK
  En sık kullanılan status koddur. servis görevi başarıyla yaptığını bu kod ile bildirir.
  
  ##### 201 Created
  Bir post metodu ile başarılı şekilde obje oluşturulduğunda verilen status kodudur. Response kısmında oluşturulan objenin bir referansı verilir.
  
  ##### 400 Bad Request
  Fonnksiyonun doğru çalışabilmesi için gerekli parametrelerde herhangi bir eksikliğin veya yanlışlığın olması durumunda kullanılır.
  
  ##### 401 Unauthorized
  Servis istekte bulunan istemcinin kimliğini doğrulayamadığında kullanılır. Login gerektiren client programlarda karşılaşılır ve istemci bu hatayı aldığında kullanıcıyı tekrar login olmaya zorlar.
  
  ##### 403 Forbidden
  Kullanıcının kimliği doğrulanabiliyor ama yetkisi bu işlem için yeterli değilse kullanılan status kodudur.
  
  ##### 404 Not Found
  Gelen isteğin bulunamaması durumunda kullanılan status koddur. Default olarak bu tepki veriliyorken özelleştirmek de mümkündür.
  
  ```
	app.Use(func(c *fiber.Ctx) error {
		response := fmt.Sprintf("%s%s is not found =(", c.BaseURL(), c.OriginalURL())
		return c.Status(404).SendString(response)
	})
  ```
⚠️ Dikkat ⚠️ bu yöntemi kullanmak istiyorsanız tüm servislerin en altına yazmalısınız. Aksi halde tüm isteklerden 404 alabilirsiniz. 👍
  
  ##### 500 Internal Server Error
  Sunucu en basit anlamda bir hata dolayısıyla çalışamazsa bu kod kullanılır. Bu tip hata kodu genellikle sunucudaki fonksiyonlardan kaynaklıdır.
  
  ##### 503 Service Unavailable
  'Sunucu geçici süreyle çevrimdışıdır. Sabrınız için teşekkürler' gibi hataları muhtemelen görmüşsünüzdür. İşte o hatanın kodu.
 
 
 
 
 
 
 
 
