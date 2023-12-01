import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 50, // Número de usuários virtuais simulados
  duration: '1m', // Duração do teste
};

export default function () {
  // Acesse apenas o endpoint principal
  let response = http.get('http://goapp.172.26.58.248.nip.io/');

  // Verifique se a solicitação foi bem-sucedida
  check(response, {
    'status is 200': (r) => r.status === 200,
  });
}

