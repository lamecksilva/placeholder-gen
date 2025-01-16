# Placeholder Generator

A simple project to generate placeholder images written in Go.

## Installation

Download with git clone and install dependencies

```bash
$ git clone git@github.com:lamecksilva/placeholder-gen.git
$ cd placeholder-gen
$ go mod download
```

I recommend running with [Air]("https://github.com/air-verse/air") package:

```bash
$ air
```

The server will start in port _8080_

## API Reference

#### Generate Placeholder

```http
  GET /generate
```

| Parameter     | Type     | Description                      |
| :------------ | :------- | :------------------------------- |
| `width`       | `number` | **Required**. Width of Image     |
| `heigth`      | `number` | **Required**. Heigth of Image    |
| `color`       | `string` | Background color                 |
| `label`       | `string` | Label of image, placed in center |
| `label-color` | `string` | Color of label text              |

OBS: The default label color is contrast color of Background color, eg:

> _background color white -> label color black._

## Authors

- [@lamecksilva](https://www.github.com/lamecksilva)
