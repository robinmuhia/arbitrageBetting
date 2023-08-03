"use client";

import { Box } from "@mui/material";
import { useState } from "react";

export default function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  return (
    <Box>
      {isLoggedIn ? (
        <></>
      ) : (
        <Box
          sx={{
            marginTop: "20px",
            display: "inline-grid",
            gridAutoFlow: "row",
            gridTemplateColumns: "repeat(2, 1fr)",
            gridTemplateRows: "repeat(1, 1fr)",
            columnGap: "50px",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Box></Box>
          <Box></Box>
        </Box>
      )}
    </Box>
  );
}
