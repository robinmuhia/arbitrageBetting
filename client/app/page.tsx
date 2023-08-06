"use client";

import { Box } from "@mui/material";
import { useState } from "react";
import FAQs from "@/components/FAQs";
import SignUp from "@/components/Signup";
import SignIn from "@/components/Signin";

export default function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [isSignUp, setIsSignUp] = useState<boolean>(true);
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
              columnGap: "80px",
              justifyContent: "center",
              alignItems: "center",
            }}
          >
            <Box>
              <FAQs />
            </Box>
            <Box>
              {isSignUp ? (
                <SignUp setIsSignUp={setIsSignUp} />
              ) : (
                <SignIn setIsLoggedIn={setIsLoggedIn} />
              )}
            </Box>
          </Box>
        </Box>
      )}
    </Box>
  );
}
