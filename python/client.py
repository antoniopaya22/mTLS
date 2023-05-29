import socket
import ssl

# ===================> Client <===================

def connect_to_tls_server(
    host: str,
    port: int,
    server_cert: str,
    client_cert: str,
    client_key: str
):
    """Connect to a TLS server.

    Args:
        host (str): The server host.
        port (int): The server port.
        server_cert (str): The Server certificate.
        client_cert (str): The client certificate.
        client_key (str): The client private key.
    
    Returns:
        None
    """
    # Create a socket and connect to the server
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.connect((host, port))

    # Create an SSL context
    context = ssl.create_default_context(ssl.Purpose.SERVER_AUTH)
    context.load_cert_chain(certfile=client_cert, keyfile=client_key)
    context.load_verify_locations(cafile=server_cert)

    # Wrap the socket with the SSL context -> Handshake
    ssl_client_socket = context.wrap_socket(client_socket, server_hostname=host)

    # Send data to the server
    data = b"Hello from the client!"
    ssl_client_socket.send(data)
    print(f'[+] Sent: {data}')
    # Receive data from the server
    encrypted_data = ssl_client_socket.recv(1024)
    data = encrypted_data.decode('utf-8')
    print(f'[+] Received: {data}')
    # Close the socket
    ssl_client_socket.close()
    
    

if __name__ == '__main__':
    connect_to_tls_server(
        host='localhost',
        port=12345,
        server_cert='../certs/server.crt',
        client_cert='../certs/client.crt',
        client_key='../certs/client.key'
    )