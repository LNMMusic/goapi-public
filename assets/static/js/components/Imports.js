const { useState, useEffect } = React;

// Metadata
const CONNECTION = "http://127.0.0.1:8080"

// Utils
async function readFetch(url='', auth='') {
    const response = await fetch(url, {credentials:"omit"});
    return response.json();
}

async function writeFetch(url='', method='POST', body={}, auth='') {
    let request = {
        method:         method,
        
        mode:           'cors',
        cache:          'no-cache',
        credentials:    'omit',

        headers:        {
            'Content-Type': 'application/json',
            // 'Content-Type': 'application/x-www-form-urlencoded'
        },
        body:           JSON.stringify(body)
    }
    if (auth) {
        request.headers['Authorization'] = `Bearer ${auth}`
    }

    const response = await fetch(url, request);
    return response.json();
}