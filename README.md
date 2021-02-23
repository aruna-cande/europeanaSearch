# europeanaSearch(In progress)
[Europeana](https://www.europeana.eu/en) works with thousands of European archives, libraries and museums to share cultural heritage for enjoyment, education and research.

This microservice enables you make search requests against europeana Api.

## Build
```bash
docker build -t europeana_search:1.0 .
```

## Run
```bash
docker run -p 8080:8080 -d --name europeana-search europeana_search:1.0
```

## Example

To get an api key you will need to create an account using the following [link](https://pro.europeana.eu/pages/get-api).

```bash
curl --location --request GET 'localhost:8080/search?query=bla&media=true&thumbnail=true&Landingpage=true&rows=10' \
--header 'Authorization: Bearer bearer_token'
```


