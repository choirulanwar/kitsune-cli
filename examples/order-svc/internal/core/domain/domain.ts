import { Field, ObjectType } from "@nestjs/graphql";
import { Order as IOrder } from "__generated__/proto/order";

@ObjectType()
export class Order implements IOrder {
  @Field(() => String)
  id?: string;

  @Field(() => String)
  short_id?: string;

  @Field(() => String)
  type?: ProductOrderType;

  @Field(() => [OrderProduct])
  products?: OrderProduct[];

  @Field(() => String)
  promo_id?: string;

  @Field(() => String)
  status?: OrderStatusType;

  @Field(() => Address, { nullable: true })
  delivery_address?: Address;

  @Field(() => String)
  user_id?: string;

  @Field(() => String)
  merchant_id?: string;

  @Field(() => String)
  payment_id?: string;

  @Field(() => Number)
  created_at?: number;

  @Field(() => Number)
  updated_at?: number;
}
