# National tax

## Make Zip to deploy function to lambda

```sh
bash deploy.sh
```

## heroku release
```sh
git push heroku master
```

## run heroku console
```sh
heroku run ./bin/national-tax
```

## manage batch
```sh
heroku addons:open scheduler
```