"use client";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuIndicator,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  NavigationMenuViewport,
} from "@/components/ui/navigation-menu";
import Link from "next/link";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card";
import authStore from "@/stores/authStore";

function Header() {
  const { account, isConnected, connectToMetaMask } = authStore();

  return (
    <header className="flex flex-row items-center justify-between al px-8 py-2">
      <h1 className="text-white text-xl font-bold">LockChain</h1>
      <div className="flex flex-row gap-10">
        <NavigationMenu>
          <NavigationMenuList className="gap-10">
            <NavigationMenuItem>
              <Link href="/" legacyBehavior passHref>
                <NavigationMenuLink>Home</NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <Link href="/add-new" legacyBehavior passHref>
                <NavigationMenuLink>Add New</NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>

        <HoverCard>
          <HoverCardTrigger className="cursor-pointer">
            <Avatar>
              <AvatarImage src="https://github.com/shadcn.png" />
              <AvatarFallback>BC</AvatarFallback>
            </Avatar>
          </HoverCardTrigger>
          <HoverCardContent className="bg-black border-slate-700 text-white">
            <div className="flex flex-row items-center gap-x-4">
              <Avatar>
                <AvatarImage src="https://github.com/shadcn.png" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
              <div>
                <p>Your current account: </p>
                <p>{account}</p>
              </div>
            </div>
          </HoverCardContent>
        </HoverCard>
      </div>
    </header>
  );
}

export default Header;
