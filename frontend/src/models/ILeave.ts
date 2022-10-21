import { DoctorInterface } from "./IDoctor";
import { EvidenceInterface } from "./IEvidence";
import { TypeInterface } from "./IType";

export interface LeaveInterface {

    ID: number,
    Reason: string;
    Fdate:Date | null;
    Ldate: Date | null;

    DoctorID?: number;
    Doctor?: DoctorInterface;

    TypeID?:number;
    Type?: TypeInterface;

    EvidenceID?:number;
    Evidence?:EvidenceInterface;
   }