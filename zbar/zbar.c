#include <stdio.h>
#include <stdlib.h>
#include <png.h>
#include <jpeglib.h>
#include <zbar.h>
#include <string.h>

zbar_image_scanner_t *scanner = NULL;

static void get_png_data(
  const char *name,
  int *width, int *height,
  void **raw
) {
  FILE *file = fopen(name, "rb");

  if(!file) exit(1);

  png_structp png = png_create_read_struct(
    PNG_LIBPNG_VER_STRING,
    NULL, NULL, NULL
  );

  if(!png) exit(1);

  if(setjmp(png_jmpbuf(png))) exit(1);

  png_infop info = png_create_info_struct(png);

  if(!info) exit(1);

  png_init_io(png, file);
  png_read_info(png, info);

  /* configure for 8bpp grayscale input */
  int color = png_get_color_type(png, info);
  int bits = png_get_bit_depth(png, info);

  if(color & PNG_COLOR_TYPE_PALETTE) {
    png_set_palette_to_rgb(png);
  }

  if(color == PNG_COLOR_TYPE_GRAY && bits < 8) {
    png_set_expand_gray_1_2_4_to_8(png);
  }

  if(bits == 16) {
    png_set_strip_16(png);
  }

  if(color & PNG_COLOR_MASK_ALPHA) {
    png_set_strip_alpha(png);
  }

  if(color & PNG_COLOR_MASK_COLOR) {
    png_set_rgb_to_gray_fixed(png, 1, -1, -1);
  }

  /* allocate image */
  *width = png_get_image_width(png, info);
  *height = png_get_image_height(png, info);
  *raw = malloc(*width * *height);

  png_bytep rows[*height];
  int i;
  for(i = 0; i < *height; i++) {
    rows[i] = *raw + (*width * i);
  }
  png_read_image(png, rows);
}

static void get_jpeg_data(
  const char *filename,
  int *width, int *height,
  void **raw
) {
  struct jpeg_decompress_struct cinfo;
  struct jpeg_error_mgr err;
  int ret;
  FILE * infile;
  JSAMPARRAY buffer;
  unsigned char * rowptr[1];
  int row_stride;     /* physical row width in output buffer */
  if((infile = fopen(filename, "rb")) == NULL) {
    fprintf(stderr, "can't open %s\n", filename);
    exit(1);
  }

  /* Step 1: allocate and initialize JPEG decompression object */

  /* We set up the normal JPEG error routines, then override error_exit. */
  cinfo.err = jpeg_std_error(&err);

  /* Now we can initialize the JPEG decompression object. */
  jpeg_create_decompress(&cinfo);

  /* Step 2: specify data source (eg, a file) */
  jpeg_stdio_src(&cinfo, infile);

  /* Step 3: read file parameters with jpeg_read_header() */
  (void) jpeg_read_header(&cinfo, TRUE);

  /* Step 4: set parameters for decompression */
  cinfo.out_color_space = JCS_GRAYSCALE;

  /* Step 5: Start decompressor */
  (void) jpeg_start_decompress(&cinfo);

  *width = cinfo.image_width;
  *height = cinfo.image_height;
  row_stride = cinfo.output_width * cinfo.output_components;
  *raw = (void *)malloc(cinfo.output_width * cinfo.output_height * 3);
  long counter = 0;

  //step 6, read the image line by line
  unsigned bpl = cinfo.output_width * cinfo.output_components;
  JSAMPROW buf = (void*)*raw;
  JSAMPARRAY line = &buf;
  for(; cinfo.output_scanline < cinfo.output_height; buf += bpl) {
    jpeg_read_scanlines(&cinfo, line, 1);
  }

  /* Step 7: Finish decompression */
  (void) jpeg_finish_decompress(&cinfo);

  /* Step 8: Release JPEG decompression object */
  /* This is an important step since it will release a good deal of memory. */
  jpeg_destroy_decompress(&cinfo);
  fclose(infile);
}

const char *get_filename_ext(const char *filename) {
  const char *dot = strrchr(filename, '.');
  if(!dot || dot == filename) return "";
  return dot + 1;
}


int main(int argc, char **argv) {
  if(argc < 2) return(1);

  /* create a reader */
  scanner = zbar_image_scanner_create();

  /* configure the reader */
  zbar_image_scanner_set_config(scanner, 0, ZBAR_CFG_ENABLE, 1);

  /* obtain image data */
  int width = 0, height = 0;
  void *raw = NULL;

  const char *fileExt = get_filename_ext(argv[1]);
  if(strcmp(fileExt, "jpg") == 0 || strcmp(fileExt, "jpeg") == 0) {
    get_jpeg_data(argv[1], &width, &height, &raw);
  } else if(strcmp(fileExt, "png") == 0) {
    get_png_data(argv[1], &width, &height, &raw);
  } else {
    fprintf(stderr, "only support .jpg,.jpeg,.png file\n");
    exit(1);
  }

  /* wrap image data */
  zbar_image_t *image = zbar_image_create();
  zbar_image_set_format(image, zbar_fourcc('Y','8','0','0'));
  zbar_image_set_size(image, width, height);
  zbar_image_set_data(image, raw, width * height, zbar_image_free_data);

  /* scan the image for barcodes */
  int n = zbar_scan_image(scanner, image);

  /* extract results */
  const zbar_symbol_t *symbol = zbar_image_first_symbol(image);
  for(; symbol; symbol = zbar_symbol_next(symbol)) {
    /* do something useful with results */
    zbar_symbol_type_t typ = zbar_symbol_get_type(symbol);
    const char *data = zbar_symbol_get_data(symbol);
    printf("%s", data);
  }

  /* clean up */
  zbar_image_destroy(image);
  zbar_image_scanner_destroy(scanner);

  return(0);
}
