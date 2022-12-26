import { Field, InputType, PartialType } from "@nestjs/graphql";
import { IsNotEmpty, IsOptional, IsString, IsNumber } from "class-validator";
import { UpdateOrderRequest } from "__generated__/proto/order";

@InputType()
export class UpdateOrderInput extends PartialType(StoreOrderInput) {}
