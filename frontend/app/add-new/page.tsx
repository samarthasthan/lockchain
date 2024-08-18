import React from "react";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";

function AddNew() {
  return (
    <div className="p-10">
      <p className="pb-4">Add new Secret/Memory</p>
      <Textarea className="bg-black" />
      <div>
        <Button className="bg-white text-black hover:bg-white mt-4">
          Add to blockchain
        </Button>
      </div>
    </div>
  );
}

export default AddNew;
