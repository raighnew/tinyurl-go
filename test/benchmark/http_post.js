import { check, sleep } from "k6";
import http from "k6/http";

export const options = {
  stages: [
    { duration: "1m", target: 200 },
    { duration: "5m", target: 200 },
    { duration: "1m", target: 500 },
    { duration: "5m", target: 500 },
    { duration: "1m", target: 0 },
  ],
};

export default function () {
  const url = "http://localhost:8080/api/generate";

  const payload = JSON.stringify({
    url: "https://cv.raighne.xyz?xxxx",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.post(url, payload, params);
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  sleep(1)
}
