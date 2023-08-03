"use client";

import { Box } from "@mui/material";
import { useState } from "react";
import FAQs from "@/components/FAQs";

export default function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  return (
    <Box>
      {isLoggedIn ? (
        <></>
      ) : (
        <Box sx={{ display: { xs: "none", lg: "block", xl: "none" } }}>
          <Box
            sx={{
              marginLeft: "30px",
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
            <Box>
              <FAQs />
            </Box>
            <Box></Box>
          </Box>
        </Box>
      )}
    </Box>
  );
}
