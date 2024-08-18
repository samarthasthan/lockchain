"use client";

import { Button } from "@/components/ui/button";
import React from "react";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useRouter } from "next/navigation";
import authStore from "@/stores/authStore";

function Auth() {
  const { connectToMetaMask } = authStore();
  const router = useRouter();

  const handleConnect = async () => {
    await connectToMetaMask();
    router.push("/");
  };

  return (
    <div className="h-screen w-screen flex justify-center items-center">
      <div className="flex flex-col max-h-36 max-w-48 gap-y-4">
        <div>
          <h1 className="font-bold text-2xl">LockChain</h1>
          <h2 className="text-xl">
            The blockchain vault for your secrets and memories
          </h2>
        </div>
        <Button
          className="bg-white text-black hover:bg-white"
          onClick={handleConnect}
        >
          Auth with MetaMask
        </Button>
        <HoverCard>
          <HoverCardTrigger className="cursor-pointer">
            @samarthasthan
          </HoverCardTrigger>
          <HoverCardContent className="bg-black border-slate-700 text-white">
            <div className="flex flex-row items-center gap-x-4">
              <Avatar>
                <AvatarImage src="https://github.com/shadcn.png" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
              <div>
                <p>@samarthasthan</p>
                <p>Follow me on Twitter.</p>
              </div>
            </div>
          </HoverCardContent>
        </HoverCard>
      </div>
    </div>
  );
}

export default Auth;
