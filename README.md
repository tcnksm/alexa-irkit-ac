# Amazon Alexa + IRKit by Golang

This is a sample implementation of custom [Alexa Smart Home Skill](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/smart-home-skill-api-reference) by Golang. Alexa recognize your voice command and turn-of/turn-on your air conditioner by sending IR signal via [IRKit](http://getirkit.com/#IRKit-Internet-API). 

See demo on [https://vimeo.com/179021210](https://vimeo.com/179021210).

## Deploy

To deploy the custom skill (AWS labmda function), use [apex](https://github.com/apex/apex),

```bash
$ apex deploy \
     -s ACCESS_TOKEN=$ACCESS_TOKEN \
     -s IR_CLIENT_KEY=$IR_CLIENT_KEY \
     -s IR_DEVICE_ID=$IR_DEVICE_ID \
     ac
```


## Author 

[Taichi Nakashima](https://github.com/tcnksm)



