# transform

transform uses primitive library to produce images with geometric primitive. Then it wraps it an interface to serve it via HTTP.

## Usage

- Run the server with `go run main.go`. 
- Head over to <http://127.0.0.1:3000>.

<img src="upload_page.png" alt="Original Image" />

- Upload an image. Here we are using a below mentioned image.

<img src="original.jpg" alt="Original Image" width="35%" />

- These are the images returned by the service.

<img src="four_modes.png" alt="Original Image" width="75%" />

In first image, we are only using circle. In second image we are using strokes; we usually need more number of shapes for image to comes out visible. In third image we just use triangle. In the fourth and the last image uses combinatoin of 1st and 3rd.
