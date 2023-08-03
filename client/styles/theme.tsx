const tokensDark = {
  //black
  secondary: {
    100: "#ffffff",
    200: "#ffffff",
    300: "#ffffff",
    400: "#ffffff",
    500: "#000000",
    600: "#ffffff",
    700: "#ffffff",
    800: "#ffffff",
    900: "#ffffff",
  },
  //blue and white
  primary: {
    100: "#ffffff",
    200: "#ffffff",
    300: "#ffffff",
    400: "#ffffff",
    500: "#ffffff",
    600: "#cccccc",
    700: "#999999",
    800: "#666666",
    900: "#333333",
  },
  //orange
  grey: {
    100: "#bb9b6a",
    200: "#f6d5ca",
    300: "#f2c0af",
    400: "#1a202c",
    500: "#e9967a",
    600: "#ba7862",
    700: "#8c5a49",
    800: "#5d3c31",
    900: "#2f1e18",
  },
};

export const themeSettings = () => {
  return {
    palette: {
      primary: {
        ...tokensDark.primary,
        main: tokensDark.primary[400],
        light: tokensDark.primary[500],
      },
      secondary: {
        ...tokensDark.secondary,
        main: tokensDark.secondary[500],
      },
      neutral: {
        ...tokensDark.grey,
        main: tokensDark.grey[100],
      },
      background: {
        default: tokensDark.primary[400],
        alt: tokensDark.primary[500],
        black: tokensDark.grey[900],
      },
    },
    typography: {
      fontFamily: ["", "sans-serif"].join(","),
      fontSize: 12,
      h1: {
        fontFamily: ["", "sans-serif"].join(","),
        fontSize: 40,
      },
      h2: {
        fontFamily: ["", "sans-serif"].join(","),
        fontSize: 32,
      },
      h3: {
        fontFamily: ["", "sans-serif"].join(","),
        fontSize: 24,
      },
      h4: {
        fontFamily: ["", "sans-serif"].join(","),
        fontSize: 20,
      },
      h5: {
        fontFamily: ["", "sans-serif"].join(","),
        fontSize: 16,
      },
      h6: {
        fontFamily: ["", "sans-serif"].join(","),
        fontSize: 14,
      },
    },
    breakpoints: {
      values: {
        xs: 0,
        sm: 480,
        md: 600,
        lg: 768,
        xl: 1536,
      },
    },
  };
};
