function Login() {
  const realm = "myrealm"; // your realm
  const clientId = "myclient"; // your client
  const keycloakUrl = "http://localhost:8084"; // your Keycloak base URL

  const login = () => {
    window.location.href =
      `${keycloakUrl}/realms/${realm}/protocol/openid-connect/auth` +
      `?client_id=${clientId}` +
      `&redirect_uri=${encodeURIComponent("http://localhost:5173/callback")}` +
      `&response_type=code&scope=openid`;
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-xl mb-4">Please log in</h1>
      <button
        onClick={login}
        className="px-4 py-2 bg-blue-600 text-white rounded-lg"
      >
        Login with Keycloak
      </button>
    </div>
  );
}

export default Login;
