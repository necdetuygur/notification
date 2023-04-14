const topic =
  window.top.location.href.indexOf("topic") > -1
    ? window.top.location.href.split("topic=")[1].split("&")[0]
    : false;

if (!topic) {
  window.top.location.href = "./index.html?topic=" + prompt("Topic:", "");
}

firebase.initializeApp({
  messagingSenderId: "411358737390",
});

const messaging = firebase.messaging();

const subscribeTokenToTopic = (token, topic) => {
  fetch("https://iid.googleapis.com/iid/v1/" + token + "/rel/topics/" + topic, {
    method: "POST",
    headers: new Headers({
      Authorization:
        "key=AAAAX8bkM-4:APA91bFjL2mYl0BQZCrvNdPtFMaMFz03GtKnT_sli-CP9EpcgFEGWTnDRaUPAnVa8zxRIzf_XBB7gH_whZ1vnWcbdPhQPsy1NuHNsNc3yoSawlQqqnp5r_E9a50MkrnqQpioyJzBeooS",
    }),
  })
    .then((response) => {
      if (response.status < 200 || response.status >= 400) {
        throw (
          "subscribeTokenToTopic: ERR: " +
          response.status +
          " - " +
          response.text()
        );
      }
    })
    .catch((error) => {
      console.error(error);

      setTimeout(() => {
        window.top.location.href = window.top.location.href;
      }, 5e3);
    });
};

const requestPermission = () => {
  Notification.requestPermission().then((permission) => {
    if (permission === "granted") {
    } else {
      console.log("Unable to Get permission to notify.");

      setTimeout(() => {
        window.top.location.href = window.top.location.href;
      }, 5e3);
    }
  });
};

messaging.onTokenRefresh(() => {
  messaging
    .getToken()
    .then((refreshedToken) => {
      subscribeTokenToTopic(refreshedToken, topic);
    })
    .catch((err) => {
      console.log("Unable to retrieve refreshed token ", err);

      setTimeout(() => {
        window.top.location.href = window.top.location.href;
      }, 5e3);
    });
});

messaging
  .getToken()
  .then((currentToken) => {
    if (currentToken) {
      subscribeTokenToTopic(currentToken, topic);
    } else {
      console.log(
        "No Instance ID token available. Request permission to generate one."
      );

      setTimeout(() => {
        window.top.location.href = window.top.location.href;
      }, 5e3);
    }
  })
  .catch((err) => {
    console.log("An error occurred while retrieving token. ", err);

    setTimeout(() => {
      window.top.location.href = window.top.location.href;
    }, 5e3);
  });

window.addEventListener("load", requestPermission);
