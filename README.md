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

## QR code
you can join the group by scanning the QR code.

![line-qr](./line-qr.png)