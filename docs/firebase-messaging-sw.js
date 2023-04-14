importScripts("https://www.gstatic.com/firebasejs/3.7.2/firebase-app.js");
importScripts("https://www.gstatic.com/firebasejs/3.7.2/firebase-messaging.js");

firebase.initializeApp({
  messagingSenderId: "411358737390",
});

const messaging = firebase.messaging();

messaging.setBackgroundMessageHandler(function (payload) {
  payload.data.data = JSON.parse(JSON.stringify(payload.data));
  return self.registration.showNotification(payload.data.title, payload.data);
});

self.addEventListener("notificationclick", function (event) {
  const target = event.notification.data.click_action || "/";
  if (event.notification.data.link) {
    clients.openWindow(event.notification.data.link);
  }
  event.notification.close();

  event.waitUntil(
    clients
      .matchAll({
        type: "window",
        includeUncontrolled: true,
      })
      .then(function (clientList) {
        for (var i = 0; i < clientList.length; i++) {
          var client = clientList[i];
          if (client.url === target && "focus" in client) {
            return client.focus();
          }
        }
        return clients.openWindow(target);
      })
  );
});
