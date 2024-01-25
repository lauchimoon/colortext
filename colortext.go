package main

import (
    "fmt"
	"image"
	"image/color"
	"strings"
    "image/png"
    "os"
    "math/rand"
)

func GetFileContents(path string) string {
    text, err := os.ReadFile(path)
    if err != nil {
        return "-1"
    }

    return string(text)
}

func MaxSlice(s []int) int {
    mx := s[0]
    for _, elem := range s {
        if elem > mx {
            mx = elem
        }
    }

    return mx
}

func GetDimensions(s string) (int, int) {
    w, h := int(0), int(0)
    textLens := []int{}

    // Determine width and height
    split := strings.Split(s, "\n")
    for _, str := range split {
        h++
        textLens = append(textLens, len(str))
    }

    w = MaxSlice(textLens)

    return w, h
}

func main() {
    args := os.Args

    if len(args) < 2 {
        fmt.Println("missing argument: file")
        os.Exit(1)
    }

    s := GetFileContents(args[1])
    w, h := GetDimensions(s)
    img := image.NewRGBA(image.Rect(0, 0, w, h))

    x, y := 0, 0

    for i := 0; i < len(s); i++ {
        c := int(s[i])
        r, g, b := uint8(rand.Intn(c)), uint8(rand.Intn(c)), uint8(rand.Intn(c))
        x++

        if c == '\n' {
            img.Set(x, y, color.RGBA{255, 255, 255, 255})
            y++
            x = 0
        } else {
            if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
                r *= 2
                g *= 2
                b *= 2
            }

            img.Set(x, y, color.RGBA{r, g, b, 255})
        }

        if y > h {
            break
        }
    }

    // Encode
    outFile, err := os.Create("out.png")
    if err != nil {
        panic(err)
    }
    defer outFile.Close()

    png.Encode(outFile, img)
}
