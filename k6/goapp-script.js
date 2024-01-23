import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 40, // Número de usuários virtuais simulados
  duration: '1m', // Duração do teste
};

export default function () {
  // Acesse apenas o endpoint principal
  let response = http.get('http://goapp.127.0.0.1.nip.io/metrics');

  // Verifique se a solicitação foi bem-sucedida
  check(response, {
    'status is 200': (r) => r.status === 200,
  });
}

