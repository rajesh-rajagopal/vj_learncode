
https://dev.infobip.com/v1/docs/api-key-create

'46ed8c0239bf4651658a85d871661a83-64d13e1e-b97f-4b02-86d4-4641428f6612'

https://dev.infobip.com/v1/docs/step-by-step-integration

{
    "applicationId": "C00EB4000DD8023343028494FD1C680D",
    "name": "MegamAfrica PIN verify",
    "configuration": {
        "pinTimeToLive": 900000,
        "pinAttempts": 10,
        "verificationAttempts": 1,
        "verificationIntervalLength": 3000,
        "initiationAttempts": 3,
        "initiationIntervalLength": 86400000,
        "overallInitiationAttempts": 10000,
        "overallInitiationIntervalLength": 86400000,
        "allowMultiplePinVerifications": true
    },
    "enabled": true,
    "processId": "7CA86F8D511EE8D4B2CAD9AE25067DC4"
}

https://dev.infobip.com/docs/message-create


{
    "messageId": "1EB03942A9025BC0F1A462810467FA60",
    "applicationId": "C00EB4000DD8023343028494FD1C680D",
    "pinPlaceholder": "<pin>",
    "messageText": "Welcome to MegamAfrica Cloud note your pin :<pin>",
    "pinLength": 4,
    "pinType": "NUMERIC",
    "language": "en"
}
