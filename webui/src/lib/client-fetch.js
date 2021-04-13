const encode = val => encodeURIComponent(val).replace(/%3A/gi, ':').replace(/%24/g, '$').replace(/%2C/gi, ',').replace(/%20/g, '+').replace(/%5B/gi, '[').replace(/%5D/gi, ']');

const buildURL = (url, params) => {
    let parts = [];

    Object.keys(params).forEach(function serialize(key) {
        let val = params[key];
        if (val === null || typeof val === 'undefined') {
            return;
        }

        if (Array.isArray(val)) {
            key = key + '[]';
        } else {
            val = [val];
        }

        val.forEach(function parseValue(v) {
            if (toString.call(v) === '[object Date]') {
                v = v.toISOString();
            } else if (v !== null && typeof v === 'object') {
                v = JSON.stringify(v);
            }
            parts.push(encode(key) + '=' + encode(v));
        });
    })

    const serializedParams = parts.join('&');

    if (serializedParams) {
        const hashmarkIndex = url.indexOf('#');
        if (hashmarkIndex !== -1) {
            url = url.slice(0, hashmarkIndex);
        }

        url += (url.indexOf('?') === -1 ? '?' : '&') + serializedParams;
    }

    return url;
};

const doRequest = async ({ url, params, ...options }) => {
    if (options.body) {
        options.body = JSON.stringify(options.body)
    }

    if (params) {
        url = buildURL(url, params);
    }

    if (options.method !== 'GET' && options.method !== 'POST') {
        options.redirect = 'manual';
    }

    const response = await fetch(url, {
        ...options,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            ...options.headers
        },
    });
    if (response.ok) {
        const text = await response.text();
        return {
            data: text.length ? JSON.parse(text) : null
        }
    } else {
        const corsDenial = ['opaque', 'opaqueredirect'].includes(response.type);
        return Promise.reject({
                message: corsDenial ? 'Please reload this page' : await response.text(),
                status: response.status,
                statusText: response.statusText,
                corsDenial
            }
        )
    }
}

const client = {
    get: (url, options) => {
        return doRequest({ url, method: 'GET', ...options });
    },
    post: (url, data, options) => {
        return doRequest({ url, method: 'POST', body: data, ...options });
    },
    put: (url, data, options) => {
        return doRequest({ url, method: 'PUT', body: data, ...options });
    },
    delete: (url, { data, ...options } = {}) => {
        return doRequest({ url, method: 'DELETE', body: data, ...options });
    }
}

export default client;
