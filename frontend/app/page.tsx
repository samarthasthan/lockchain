"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import Header from "@/components/header/header";
import authStore from "@/stores/authStore";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export default function Home() {
  const { account, isConnected, connectToMetaMask } = authStore();
  const router = useRouter();

  useEffect(() => {
    if (!isConnected) {
      connectToMetaMask().then(() => {
        if (!isConnected) {
          router.push("/auth");
        }
      });
    }
  }, [isConnected, connectToMetaMask, router]);

  if (!isConnected) {
    return null; // Render nothing while checking
  }

  const data = [
    {
      id: "1",
      content: "My first secret content",
      timestamp: "2024-08-18T12:00:00Z",
    },
    {
      id: "2",
      content: "Another secret that I need to keep safe",
      timestamp: "2024-08-18T12:30:00Z",
    },
    {
      id: "3",
      content: "Password for my online account",
      timestamp: "2024-08-18T13:00:00Z",
    },
    {
      id: "4",
      content: "Confidential project details",
      timestamp: "2024-08-18T13:30:00Z",
    },
    {
      id: "5",
      content: "Personal note about a future plan",
      timestamp: "2024-08-18T14:00:00Z",
    },
    {
      id: "6",
      content: "Secret recipe for a special dish",
      timestamp: "2024-08-18T14:30:00Z",
    },
    {
      id: "7",
      content: "Hidden details for a surprise party",
      timestamp: "2024-08-18T15:00:00Z",
    },
    {
      id: "8",
      content: "Secure storage information",
      timestamp: "2024-08-18T15:30:00Z",
    },
    {
      id: "9",
      content: "Backup code for a critical system",
      timestamp: "2024-08-18T16:00:00Z",
    },
    {
      id: "10",
      content: "Private conversation notes",
      timestamp: "2024-08-18T16:30:00Z",
    },
    {
      id: "11",
      content: "Sensitive financial information",
      timestamp: "2024-08-18T17:00:00Z",
    },
    {
      id: "12",
      content: "Important contact details",
      timestamp: "2024-08-18T17:30:00Z",
    },
    {
      id: "13",
      content: "Details about a personal project",
      timestamp: "2024-08-18T18:00:00Z",
    },
    {
      id: "14",
      content: "Notes on a confidential research",
      timestamp: "2024-08-18T18:30:00Z",
    },
    {
      id: "15",
      content: "Private meeting agenda",
      timestamp: "2024-08-18T19:00:00Z",
    },
    {
      id: "16",
      content: "Access codes for a secured area",
      timestamp: "2024-08-18T19:30:00Z",
    },
    {
      id: "17",
      content: "Special instructions for a task",
      timestamp: "2024-08-18T20:00:00Z",
    },
    {
      id: "18",
      content: "Confidential health information",
      timestamp: "2024-08-18T20:30:00Z",
    },
    {
      id: "19",
      content: "Encrypted message for a friend",
      timestamp: "2024-08-18T21:00:00Z",
    },
    {
      id: "20",
      content: "Sensitive legal document details",
      timestamp: "2024-08-18T21:30:00Z",
    },
  ];

  return (
    <main>
      <div className="px-10 pt-5">
        <h2>List</h2>
        <div className="mt-6">
          {data.map((item, key) => (
            <Card className="bg-black mb-10" key={key}>
              <CardHeader>
                <CardTitle>{item.content}</CardTitle>
                {/* <CardDescription>{item.content}</CardDescription> */}
              </CardHeader>
              {/* <CardContent>
              <p>Card Content</p>
            </CardContent>
            <CardFooter>
              <p>Card Footer</p>
            </CardFooter> */}
            </Card>
          ))}
        </div>
      </div>
    </main>
  );
}
