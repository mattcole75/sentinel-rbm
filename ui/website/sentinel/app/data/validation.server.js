function isValidName(value) {
    return value && value.trim().length > 0 && value.trim().length <= 64;
}

function isValidDescription(value) {
    return value && value.trim().length > 0 && value.trim().length <= 512;
}

function isValidTitle(value) {
    return value && value.trim().length > 0 && value.trim().length <= 64;
}

function isValidStatement(value) {
    return value && value.trim().length > 0 && value.trim().length <= 512;
}

function isValidReference(value) {
    return value && value.trim().length > 0 && value.trim().length <= 64;
}

function isValidSource(value) {
    return value && value.trim().length > 0 && value.trim().length <= 64;
}

function isValidEmail(value) {
    return String(value)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
}

function isValidPassword(value) {
    return value && value.trim().length >= 8;
}
// function isValidPassword(password) {
//     /* giving credit... http://www.thegeekstuff.com/2008/06/the-ultimate-guide-for-creating-strong-passwords */
//     var ucAlphaCount = 0, lcAlphaCount = 0, numberCount = 0, specialCount = 0,
//         parsedPwd = password.split(''), specialChars = ['!', '@', '#', '$', '%', '&', '_'];

//     //0. 8 characters in length
//     if (password.toString().trim().length < 8) { return false; }
//     //1. at least one lower case alphabet
//     //2. at least one upper case alphabet
//     //3. at least one number
//     //4. at least one special character
//     parsedPwd.forEach(function (character) {
//         if (character === character.toUpperCase()) { ucAlphaCount++; }
//         if (character === character.toLowerCase()) { lcAlphaCount++; }
//         if (!isNaN(character)) { numberCount++; }
//         if (specialChars.indexOf(character) > -1) { specialCount++; }
//     });
//     return (ucAlphaCount > 0 && lcAlphaCount > 0 && numberCount > 0 && specialCount > 0);
// }
  
// function isValidAmount(value) {
//     const amount = parseFloat(value);
//     return !isNaN(amount) && amount > 0;
// }
  
// function isValidDate(value) {
//     return value && new Date(value).getTime() < new Date().getTime();
// }

export function validateCredentials(input) {
    let validationErrors = {};

    if (!isValidEmail(input.email)) {
        validationErrors.email = "Invalid email";
    }

    if(!isValidPassword(input.password)) {
        validationErrors.password = "Invalid password";
    }


    if(Object.keys(validationErrors).length > 0) {
        console.log("mca", validationErrors);
        throw validationErrors;
    }
}

export function validateSystemInput(input) {
    let validationErrors = {};
    
    if (!isValidName (input.name)) {
        validationErrors.name = "Invalid name. Must be at most 64 characters long";
    }

    if (!isValidDescription(input.description)) {
        validationErrors.description = "Invalid description. Must be at most 512 characters long";
    }

    if (Object.keys(validationErrors).length > 0) {
        throw validationErrors;
    }
}
  
export function validateRequirementInput(input) {
    let validationErrors = {};
  
    if (!isValidTitle(input.title)) {
        validationErrors.title = "Invalid title. Must be at most 64 characters long";
    }

    if (!isValidStatement(input.statement)) {
        validationErrors.statement = "Invalid statement. Must be at most 512 characters long";
    }

    if (!isValidReference(input.reference)) {
        validationErrors.reference = "Invalid reference. Must be at most 64 characters long";
    }

    if (!isValidSource(input.referenceSource)) {
        validationErrors.referenceSource = "Invalid source. Must be at most 64 characters long";
    }
  
    // if (!isValidAmount(input.amount)) {
    //     validationErrors.amount = 'Invalid amount. Must be a number greater than zero.'
    // }
  
    // if (!isValidDate(input.date)) {
    //     validationErrors.date = 'Invalid date. Must be a date before today.'
    // }
  
    if (Object.keys(validationErrors).length > 0) {
        throw validationErrors;
    }
}
