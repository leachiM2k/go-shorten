const refreshTokenSetup = res => {
    let refreshTiming = (res.tokenObj.expires_in || 3600 - 5 * 60) * 1000;

    const refreshToken = async () => {
        const newAuthRes = await res.reloadAuthResponse();
        refreshTiming = (newAuthRes.expires_in || 3600 - 5 * 60) * 1000;

        console.log('*************************** newAuthRes', newAuthRes);
        // saveUser(newAuthRes.access_token); // <-- save new token
        console.log('*************************** new auth token', newAuthRes.id_token);

        // Setup the other timer after the first one
        setTimeout(refreshToken, refreshTiming);
    };

    // Setup the first refresh timer
    setTimeout(refreshToken, refreshTiming);
};

export default refreshTokenSetup;
