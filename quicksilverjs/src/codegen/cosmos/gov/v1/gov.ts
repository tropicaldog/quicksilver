import { Coin, CoinAmino, CoinSDKType } from "../../base/v1beta1/coin";
import { Any, AnyProtoMsg, AnyAmino, AnySDKType } from "../../../google/protobuf/any";
import { Timestamp, TimestampAmino, TimestampSDKType } from "../../../google/protobuf/timestamp";
import { Duration, DurationAmino, DurationSDKType } from "../../../google/protobuf/duration";
import { Long, isSet, DeepPartial, toTimestamp, fromTimestamp } from "../../../helpers";
import * as _m0 from "protobufjs/minimal";
export const protobufPackage = "cosmos.gov.v1";
/** VoteOption enumerates the valid vote options for a given governance proposal. */
export enum VoteOption {
  /** VOTE_OPTION_UNSPECIFIED - VOTE_OPTION_UNSPECIFIED defines a no-op vote option. */
  VOTE_OPTION_UNSPECIFIED = 0,
  /** VOTE_OPTION_YES - VOTE_OPTION_YES defines a yes vote option. */
  VOTE_OPTION_YES = 1,
  /** VOTE_OPTION_ABSTAIN - VOTE_OPTION_ABSTAIN defines an abstain vote option. */
  VOTE_OPTION_ABSTAIN = 2,
  /** VOTE_OPTION_NO - VOTE_OPTION_NO defines a no vote option. */
  VOTE_OPTION_NO = 3,
  /** VOTE_OPTION_NO_WITH_VETO - VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option. */
  VOTE_OPTION_NO_WITH_VETO = 4,
  UNRECOGNIZED = -1,
}
export const VoteOptionSDKType = VoteOption;
export const VoteOptionAmino = VoteOption;
export function voteOptionFromJSON(object: any): VoteOption {
  switch (object) {
    case 0:
    case "VOTE_OPTION_UNSPECIFIED":
      return VoteOption.VOTE_OPTION_UNSPECIFIED;
    case 1:
    case "VOTE_OPTION_YES":
      return VoteOption.VOTE_OPTION_YES;
    case 2:
    case "VOTE_OPTION_ABSTAIN":
      return VoteOption.VOTE_OPTION_ABSTAIN;
    case 3:
    case "VOTE_OPTION_NO":
      return VoteOption.VOTE_OPTION_NO;
    case 4:
    case "VOTE_OPTION_NO_WITH_VETO":
      return VoteOption.VOTE_OPTION_NO_WITH_VETO;
    case -1:
    case "UNRECOGNIZED":
    default:
      return VoteOption.UNRECOGNIZED;
  }
}
export function voteOptionToJSON(object: VoteOption): string {
  switch (object) {
    case VoteOption.VOTE_OPTION_UNSPECIFIED:
      return "VOTE_OPTION_UNSPECIFIED";
    case VoteOption.VOTE_OPTION_YES:
      return "VOTE_OPTION_YES";
    case VoteOption.VOTE_OPTION_ABSTAIN:
      return "VOTE_OPTION_ABSTAIN";
    case VoteOption.VOTE_OPTION_NO:
      return "VOTE_OPTION_NO";
    case VoteOption.VOTE_OPTION_NO_WITH_VETO:
      return "VOTE_OPTION_NO_WITH_VETO";
    case VoteOption.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
/** ProposalStatus enumerates the valid statuses of a proposal. */
export enum ProposalStatus {
  /** PROPOSAL_STATUS_UNSPECIFIED - PROPOSAL_STATUS_UNSPECIFIED defines the default propopsal status. */
  PROPOSAL_STATUS_UNSPECIFIED = 0,
  /**
   * PROPOSAL_STATUS_DEPOSIT_PERIOD - PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
   * period.
   */
  PROPOSAL_STATUS_DEPOSIT_PERIOD = 1,
  /**
   * PROPOSAL_STATUS_VOTING_PERIOD - PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
   * period.
   */
  PROPOSAL_STATUS_VOTING_PERIOD = 2,
  /**
   * PROPOSAL_STATUS_PASSED - PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
   * passed.
   */
  PROPOSAL_STATUS_PASSED = 3,
  /**
   * PROPOSAL_STATUS_REJECTED - PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
   * been rejected.
   */
  PROPOSAL_STATUS_REJECTED = 4,
  /**
   * PROPOSAL_STATUS_FAILED - PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
   * failed.
   */
  PROPOSAL_STATUS_FAILED = 5,
  UNRECOGNIZED = -1,
}
export const ProposalStatusSDKType = ProposalStatus;
export const ProposalStatusAmino = ProposalStatus;
export function proposalStatusFromJSON(object: any): ProposalStatus {
  switch (object) {
    case 0:
    case "PROPOSAL_STATUS_UNSPECIFIED":
      return ProposalStatus.PROPOSAL_STATUS_UNSPECIFIED;
    case 1:
    case "PROPOSAL_STATUS_DEPOSIT_PERIOD":
      return ProposalStatus.PROPOSAL_STATUS_DEPOSIT_PERIOD;
    case 2:
    case "PROPOSAL_STATUS_VOTING_PERIOD":
      return ProposalStatus.PROPOSAL_STATUS_VOTING_PERIOD;
    case 3:
    case "PROPOSAL_STATUS_PASSED":
      return ProposalStatus.PROPOSAL_STATUS_PASSED;
    case 4:
    case "PROPOSAL_STATUS_REJECTED":
      return ProposalStatus.PROPOSAL_STATUS_REJECTED;
    case 5:
    case "PROPOSAL_STATUS_FAILED":
      return ProposalStatus.PROPOSAL_STATUS_FAILED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ProposalStatus.UNRECOGNIZED;
  }
}
export function proposalStatusToJSON(object: ProposalStatus): string {
  switch (object) {
    case ProposalStatus.PROPOSAL_STATUS_UNSPECIFIED:
      return "PROPOSAL_STATUS_UNSPECIFIED";
    case ProposalStatus.PROPOSAL_STATUS_DEPOSIT_PERIOD:
      return "PROPOSAL_STATUS_DEPOSIT_PERIOD";
    case ProposalStatus.PROPOSAL_STATUS_VOTING_PERIOD:
      return "PROPOSAL_STATUS_VOTING_PERIOD";
    case ProposalStatus.PROPOSAL_STATUS_PASSED:
      return "PROPOSAL_STATUS_PASSED";
    case ProposalStatus.PROPOSAL_STATUS_REJECTED:
      return "PROPOSAL_STATUS_REJECTED";
    case ProposalStatus.PROPOSAL_STATUS_FAILED:
      return "PROPOSAL_STATUS_FAILED";
    case ProposalStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
/** WeightedVoteOption defines a unit of vote for vote split. */
export interface WeightedVoteOption {
  option: VoteOption;
  weight: string;
}
export interface WeightedVoteOptionProtoMsg {
  typeUrl: "/cosmos.gov.v1.WeightedVoteOption";
  value: Uint8Array;
}
/** WeightedVoteOption defines a unit of vote for vote split. */
export interface WeightedVoteOptionAmino {
  option: VoteOption;
  weight: string;
}
export interface WeightedVoteOptionAminoMsg {
  type: "cosmos-sdk/v1/WeightedVoteOption";
  value: WeightedVoteOptionAmino;
}
/** WeightedVoteOption defines a unit of vote for vote split. */
export interface WeightedVoteOptionSDKType {
  option: VoteOption;
  weight: string;
}
/**
 * Deposit defines an amount deposited by an account address to an active
 * proposal.
 */
export interface Deposit {
  proposalId: Long;
  depositor: string;
  amount: Coin[];
}
export interface DepositProtoMsg {
  typeUrl: "/cosmos.gov.v1.Deposit";
  value: Uint8Array;
}
/**
 * Deposit defines an amount deposited by an account address to an active
 * proposal.
 */
export interface DepositAmino {
  proposal_id: string;
  depositor: string;
  amount: CoinAmino[];
}
export interface DepositAminoMsg {
  type: "cosmos-sdk/v1/Deposit";
  value: DepositAmino;
}
/**
 * Deposit defines an amount deposited by an account address to an active
 * proposal.
 */
export interface DepositSDKType {
  proposal_id: Long;
  depositor: string;
  amount: CoinSDKType[];
}
/** Proposal defines the core field members of a governance proposal. */
export interface Proposal {
  id: Long;
  messages: Any[];
  status: ProposalStatus;
  /**
   * final_tally_result is the final tally result of the proposal. When
   * querying a proposal via gRPC, this field is not populated until the
   * proposal's voting period has ended.
   */
  finalTallyResult: TallyResult;
  submitTime: Date;
  depositEndTime: Date;
  totalDeposit: Coin[];
  votingStartTime: Date;
  votingEndTime: Date;
  /** metadata is any arbitrary metadata attached to the proposal. */
  metadata: string;
}
export interface ProposalProtoMsg {
  typeUrl: "/cosmos.gov.v1.Proposal";
  value: Uint8Array;
}
/** Proposal defines the core field members of a governance proposal. */
export interface ProposalAmino {
  id: string;
  messages: AnyAmino[];
  status: ProposalStatus;
  /**
   * final_tally_result is the final tally result of the proposal. When
   * querying a proposal via gRPC, this field is not populated until the
   * proposal's voting period has ended.
   */
  final_tally_result?: TallyResultAmino;
  submit_time?: Date;
  deposit_end_time?: Date;
  total_deposit: CoinAmino[];
  voting_start_time?: Date;
  voting_end_time?: Date;
  /** metadata is any arbitrary metadata attached to the proposal. */
  metadata: string;
}
export interface ProposalAminoMsg {
  type: "cosmos-sdk/v1/Proposal";
  value: ProposalAmino;
}
/** Proposal defines the core field members of a governance proposal. */
export interface ProposalSDKType {
  id: Long;
  messages: AnySDKType[];
  status: ProposalStatus;
  final_tally_result: TallyResultSDKType;
  submit_time: Date;
  deposit_end_time: Date;
  total_deposit: CoinSDKType[];
  voting_start_time: Date;
  voting_end_time: Date;
  metadata: string;
}
/** TallyResult defines a standard tally for a governance proposal. */
export interface TallyResult {
  yesCount: string;
  abstainCount: string;
  noCount: string;
  noWithVetoCount: string;
}
export interface TallyResultProtoMsg {
  typeUrl: "/cosmos.gov.v1.TallyResult";
  value: Uint8Array;
}
/** TallyResult defines a standard tally for a governance proposal. */
export interface TallyResultAmino {
  yes_count: string;
  abstain_count: string;
  no_count: string;
  no_with_veto_count: string;
}
export interface TallyResultAminoMsg {
  type: "cosmos-sdk/v1/TallyResult";
  value: TallyResultAmino;
}
/** TallyResult defines a standard tally for a governance proposal. */
export interface TallyResultSDKType {
  yes_count: string;
  abstain_count: string;
  no_count: string;
  no_with_veto_count: string;
}
/**
 * Vote defines a vote on a governance proposal.
 * A Vote consists of a proposal ID, the voter, and the vote option.
 */
export interface Vote {
  proposalId: Long;
  voter: string;
  options: WeightedVoteOption[];
  /** metadata is any  arbitrary metadata to attached to the vote. */
  metadata: string;
}
export interface VoteProtoMsg {
  typeUrl: "/cosmos.gov.v1.Vote";
  value: Uint8Array;
}
/**
 * Vote defines a vote on a governance proposal.
 * A Vote consists of a proposal ID, the voter, and the vote option.
 */
export interface VoteAmino {
  proposal_id: string;
  voter: string;
  options: WeightedVoteOptionAmino[];
  /** metadata is any  arbitrary metadata to attached to the vote. */
  metadata: string;
}
export interface VoteAminoMsg {
  type: "cosmos-sdk/v1/Vote";
  value: VoteAmino;
}
/**
 * Vote defines a vote on a governance proposal.
 * A Vote consists of a proposal ID, the voter, and the vote option.
 */
export interface VoteSDKType {
  proposal_id: Long;
  voter: string;
  options: WeightedVoteOptionSDKType[];
  metadata: string;
}
/** DepositParams defines the params for deposits on governance proposals. */
export interface DepositParams {
  /** Minimum deposit for a proposal to enter voting period. */
  minDeposit: Coin[];
  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   *  months.
   */
  maxDepositPeriod: Duration;
}
export interface DepositParamsProtoMsg {
  typeUrl: "/cosmos.gov.v1.DepositParams";
  value: Uint8Array;
}
/** DepositParams defines the params for deposits on governance proposals. */
export interface DepositParamsAmino {
  /** Minimum deposit for a proposal to enter voting period. */
  min_deposit: CoinAmino[];
  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   *  months.
   */
  max_deposit_period?: DurationAmino;
}
export interface DepositParamsAminoMsg {
  type: "cosmos-sdk/v1/DepositParams";
  value: DepositParamsAmino;
}
/** DepositParams defines the params for deposits on governance proposals. */
export interface DepositParamsSDKType {
  min_deposit: CoinSDKType[];
  max_deposit_period: DurationSDKType;
}
/** VotingParams defines the params for voting on governance proposals. */
export interface VotingParams {
  /** Length of the voting period. */
  votingPeriod: Duration;
}
export interface VotingParamsProtoMsg {
  typeUrl: "/cosmos.gov.v1.VotingParams";
  value: Uint8Array;
}
/** VotingParams defines the params for voting on governance proposals. */
export interface VotingParamsAmino {
  /** Length of the voting period. */
  voting_period?: DurationAmino;
}
export interface VotingParamsAminoMsg {
  type: "cosmos-sdk/v1/VotingParams";
  value: VotingParamsAmino;
}
/** VotingParams defines the params for voting on governance proposals. */
export interface VotingParamsSDKType {
  voting_period: DurationSDKType;
}
/** TallyParams defines the params for tallying votes on governance proposals. */
export interface TallyParams {
  /**
   * Minimum percentage of total stake needed to vote for a result to be
   *  considered valid.
   */
  quorum: string;
  /** Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. */
  threshold: string;
  /**
   * Minimum value of Veto votes to Total votes ratio for proposal to be
   *  vetoed. Default value: 1/3.
   */
  vetoThreshold: string;
}
export interface TallyParamsProtoMsg {
  typeUrl: "/cosmos.gov.v1.TallyParams";
  value: Uint8Array;
}
/** TallyParams defines the params for tallying votes on governance proposals. */
export interface TallyParamsAmino {
  /**
   * Minimum percentage of total stake needed to vote for a result to be
   *  considered valid.
   */
  quorum: string;
  /** Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. */
  threshold: string;
  /**
   * Minimum value of Veto votes to Total votes ratio for proposal to be
   *  vetoed. Default value: 1/3.
   */
  veto_threshold: string;
}
export interface TallyParamsAminoMsg {
  type: "cosmos-sdk/v1/TallyParams";
  value: TallyParamsAmino;
}
/** TallyParams defines the params for tallying votes on governance proposals. */
export interface TallyParamsSDKType {
  quorum: string;
  threshold: string;
  veto_threshold: string;
}
function createBaseWeightedVoteOption(): WeightedVoteOption {
  return {
    option: 0,
    weight: ""
  };
}
export const WeightedVoteOption = {
  typeUrl: "/cosmos.gov.v1.WeightedVoteOption",
  aminoType: "cosmos-sdk/v1/WeightedVoteOption",
  encode(message: WeightedVoteOption, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.option !== 0) {
      writer.uint32(8).int32(message.option);
    }
    if (message.weight !== "") {
      writer.uint32(18).string(message.weight);
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): WeightedVoteOption {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWeightedVoteOption();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.option = (reader.int32() as any);
          break;
        case 2:
          message.weight = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): WeightedVoteOption {
    const obj = createBaseWeightedVoteOption();
    if (isSet(object.option)) obj.option = voteOptionFromJSON(object.option);
    if (isSet(object.weight)) obj.weight = String(object.weight);
    return obj;
  },
  toJSON(message: WeightedVoteOption): unknown {
    const obj: any = {};
    message.option !== undefined && (obj.option = voteOptionToJSON(message.option));
    message.weight !== undefined && (obj.weight = message.weight);
    return obj;
  },
  fromPartial(object: DeepPartial<WeightedVoteOption>): WeightedVoteOption {
    const message = createBaseWeightedVoteOption();
    message.option = object.option ?? 0;
    message.weight = object.weight ?? "";
    return message;
  },
  fromSDK(object: WeightedVoteOptionSDKType): WeightedVoteOption {
    return {
      option: isSet(object.option) ? voteOptionFromJSON(object.option) : -1,
      weight: object?.weight
    };
  },
  toSDK(message: WeightedVoteOption): WeightedVoteOptionSDKType {
    const obj: any = {};
    message.option !== undefined && (obj.option = voteOptionToJSON(message.option));
    obj.weight = message.weight;
    return obj;
  },
  fromAmino(object: WeightedVoteOptionAmino): WeightedVoteOption {
    return {
      option: isSet(object.option) ? voteOptionFromJSON(object.option) : -1,
      weight: object.weight
    };
  },
  toAmino(message: WeightedVoteOption): WeightedVoteOptionAmino {
    const obj: any = {};
    obj.option = message.option;
    obj.weight = message.weight;
    return obj;
  },
  fromAminoMsg(object: WeightedVoteOptionAminoMsg): WeightedVoteOption {
    return WeightedVoteOption.fromAmino(object.value);
  },
  toAminoMsg(message: WeightedVoteOption): WeightedVoteOptionAminoMsg {
    return {
      type: "cosmos-sdk/v1/WeightedVoteOption",
      value: WeightedVoteOption.toAmino(message)
    };
  },
  fromProtoMsg(message: WeightedVoteOptionProtoMsg): WeightedVoteOption {
    return WeightedVoteOption.decode(message.value);
  },
  toProto(message: WeightedVoteOption): Uint8Array {
    return WeightedVoteOption.encode(message).finish();
  },
  toProtoMsg(message: WeightedVoteOption): WeightedVoteOptionProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.WeightedVoteOption",
      value: WeightedVoteOption.encode(message).finish()
    };
  }
};
function createBaseDeposit(): Deposit {
  return {
    proposalId: Long.UZERO,
    depositor: "",
    amount: []
  };
}
export const Deposit = {
  typeUrl: "/cosmos.gov.v1.Deposit",
  aminoType: "cosmos-sdk/v1/Deposit",
  encode(message: Deposit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.proposalId.isZero()) {
      writer.uint32(8).uint64(message.proposalId);
    }
    if (message.depositor !== "") {
      writer.uint32(18).string(message.depositor);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): Deposit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeposit();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposalId = (reader.uint64() as Long);
          break;
        case 2:
          message.depositor = reader.string();
          break;
        case 3:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Deposit {
    const obj = createBaseDeposit();
    if (isSet(object.proposalId)) obj.proposalId = Long.fromValue(object.proposalId);
    if (isSet(object.depositor)) obj.depositor = String(object.depositor);
    if (Array.isArray(object?.amount)) obj.amount = object.amount.map((e: any) => Coin.fromJSON(e));
    return obj;
  },
  toJSON(message: Deposit): unknown {
    const obj: any = {};
    message.proposalId !== undefined && (obj.proposalId = (message.proposalId || Long.UZERO).toString());
    message.depositor !== undefined && (obj.depositor = message.depositor);
    if (message.amount) {
      obj.amount = message.amount.map(e => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },
  fromPartial(object: DeepPartial<Deposit>): Deposit {
    const message = createBaseDeposit();
    if (object.proposalId !== undefined && object.proposalId !== null) {
      message.proposalId = Long.fromValue(object.proposalId);
    }
    message.depositor = object.depositor ?? "";
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    return message;
  },
  fromSDK(object: DepositSDKType): Deposit {
    return {
      proposalId: object?.proposal_id,
      depositor: object?.depositor,
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromSDK(e)) : []
    };
  },
  toSDK(message: Deposit): DepositSDKType {
    const obj: any = {};
    obj.proposal_id = message.proposalId;
    obj.depositor = message.depositor;
    if (message.amount) {
      obj.amount = message.amount.map(e => e ? Coin.toSDK(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },
  fromAmino(object: DepositAmino): Deposit {
    return {
      proposalId: Long.fromString(object.proposal_id),
      depositor: object.depositor,
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromAmino(e)) : []
    };
  },
  toAmino(message: Deposit): DepositAmino {
    const obj: any = {};
    obj.proposal_id = message.proposalId ? message.proposalId.toString() : undefined;
    obj.depositor = message.depositor;
    if (message.amount) {
      obj.amount = message.amount.map(e => e ? Coin.toAmino(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },
  fromAminoMsg(object: DepositAminoMsg): Deposit {
    return Deposit.fromAmino(object.value);
  },
  toAminoMsg(message: Deposit): DepositAminoMsg {
    return {
      type: "cosmos-sdk/v1/Deposit",
      value: Deposit.toAmino(message)
    };
  },
  fromProtoMsg(message: DepositProtoMsg): Deposit {
    return Deposit.decode(message.value);
  },
  toProto(message: Deposit): Uint8Array {
    return Deposit.encode(message).finish();
  },
  toProtoMsg(message: Deposit): DepositProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.Deposit",
      value: Deposit.encode(message).finish()
    };
  }
};
function createBaseProposal(): Proposal {
  return {
    id: Long.UZERO,
    messages: [],
    status: 0,
    finalTallyResult: TallyResult.fromPartial({}),
    submitTime: new Date(),
    depositEndTime: new Date(),
    totalDeposit: [],
    votingStartTime: new Date(),
    votingEndTime: new Date(),
    metadata: ""
  };
}
export const Proposal = {
  typeUrl: "/cosmos.gov.v1.Proposal",
  aminoType: "cosmos-sdk/v1/Proposal",
  encode(message: Proposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.id.isZero()) {
      writer.uint32(8).uint64(message.id);
    }
    for (const v of message.messages) {
      Any.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(24).int32(message.status);
    }
    if (message.finalTallyResult !== undefined) {
      TallyResult.encode(message.finalTallyResult, writer.uint32(34).fork()).ldelim();
    }
    if (message.submitTime !== undefined) {
      Timestamp.encode(toTimestamp(message.submitTime), writer.uint32(42).fork()).ldelim();
    }
    if (message.depositEndTime !== undefined) {
      Timestamp.encode(toTimestamp(message.depositEndTime), writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.totalDeposit) {
      Coin.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.votingStartTime !== undefined) {
      Timestamp.encode(toTimestamp(message.votingStartTime), writer.uint32(66).fork()).ldelim();
    }
    if (message.votingEndTime !== undefined) {
      Timestamp.encode(toTimestamp(message.votingEndTime), writer.uint32(74).fork()).ldelim();
    }
    if (message.metadata !== "") {
      writer.uint32(82).string(message.metadata);
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): Proposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = (reader.uint64() as Long);
          break;
        case 2:
          message.messages.push(Any.decode(reader, reader.uint32()));
          break;
        case 3:
          message.status = (reader.int32() as any);
          break;
        case 4:
          message.finalTallyResult = TallyResult.decode(reader, reader.uint32());
          break;
        case 5:
          message.submitTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 6:
          message.depositEndTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 7:
          message.totalDeposit.push(Coin.decode(reader, reader.uint32()));
          break;
        case 8:
          message.votingStartTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 9:
          message.votingEndTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 10:
          message.metadata = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Proposal {
    const obj = createBaseProposal();
    if (isSet(object.id)) obj.id = Long.fromValue(object.id);
    if (Array.isArray(object?.messages)) obj.messages = object.messages.map((e: any) => Any.fromJSON(e));
    if (isSet(object.status)) obj.status = proposalStatusFromJSON(object.status);
    if (isSet(object.finalTallyResult)) obj.finalTallyResult = TallyResult.fromJSON(object.finalTallyResult);
    if (isSet(object.submitTime)) obj.submitTime = new Date(object.submitTime);
    if (isSet(object.depositEndTime)) obj.depositEndTime = new Date(object.depositEndTime);
    if (Array.isArray(object?.totalDeposit)) obj.totalDeposit = object.totalDeposit.map((e: any) => Coin.fromJSON(e));
    if (isSet(object.votingStartTime)) obj.votingStartTime = new Date(object.votingStartTime);
    if (isSet(object.votingEndTime)) obj.votingEndTime = new Date(object.votingEndTime);
    if (isSet(object.metadata)) obj.metadata = String(object.metadata);
    return obj;
  },
  toJSON(message: Proposal): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = (message.id || Long.UZERO).toString());
    if (message.messages) {
      obj.messages = message.messages.map(e => e ? Any.toJSON(e) : undefined);
    } else {
      obj.messages = [];
    }
    message.status !== undefined && (obj.status = proposalStatusToJSON(message.status));
    message.finalTallyResult !== undefined && (obj.finalTallyResult = message.finalTallyResult ? TallyResult.toJSON(message.finalTallyResult) : undefined);
    message.submitTime !== undefined && (obj.submitTime = message.submitTime.toISOString());
    message.depositEndTime !== undefined && (obj.depositEndTime = message.depositEndTime.toISOString());
    if (message.totalDeposit) {
      obj.totalDeposit = message.totalDeposit.map(e => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.totalDeposit = [];
    }
    message.votingStartTime !== undefined && (obj.votingStartTime = message.votingStartTime.toISOString());
    message.votingEndTime !== undefined && (obj.votingEndTime = message.votingEndTime.toISOString());
    message.metadata !== undefined && (obj.metadata = message.metadata);
    return obj;
  },
  fromPartial(object: DeepPartial<Proposal>): Proposal {
    const message = createBaseProposal();
    if (object.id !== undefined && object.id !== null) {
      message.id = Long.fromValue(object.id);
    }
    message.messages = object.messages?.map(e => Any.fromPartial(e)) || [];
    message.status = object.status ?? 0;
    if (object.finalTallyResult !== undefined && object.finalTallyResult !== null) {
      message.finalTallyResult = TallyResult.fromPartial(object.finalTallyResult);
    }
    message.submitTime = object.submitTime ?? undefined;
    message.depositEndTime = object.depositEndTime ?? undefined;
    message.totalDeposit = object.totalDeposit?.map(e => Coin.fromPartial(e)) || [];
    message.votingStartTime = object.votingStartTime ?? undefined;
    message.votingEndTime = object.votingEndTime ?? undefined;
    message.metadata = object.metadata ?? "";
    return message;
  },
  fromSDK(object: ProposalSDKType): Proposal {
    return {
      id: object?.id,
      messages: Array.isArray(object?.messages) ? object.messages.map((e: any) => Any.fromSDK(e)) : [],
      status: isSet(object.status) ? proposalStatusFromJSON(object.status) : -1,
      finalTallyResult: object.final_tally_result ? TallyResult.fromSDK(object.final_tally_result) : undefined,
      submitTime: object.submit_time ?? undefined,
      depositEndTime: object.deposit_end_time ?? undefined,
      totalDeposit: Array.isArray(object?.total_deposit) ? object.total_deposit.map((e: any) => Coin.fromSDK(e)) : [],
      votingStartTime: object.voting_start_time ?? undefined,
      votingEndTime: object.voting_end_time ?? undefined,
      metadata: object?.metadata
    };
  },
  toSDK(message: Proposal): ProposalSDKType {
    const obj: any = {};
    obj.id = message.id;
    if (message.messages) {
      obj.messages = message.messages.map(e => e ? Any.toSDK(e) : undefined);
    } else {
      obj.messages = [];
    }
    message.status !== undefined && (obj.status = proposalStatusToJSON(message.status));
    message.finalTallyResult !== undefined && (obj.final_tally_result = message.finalTallyResult ? TallyResult.toSDK(message.finalTallyResult) : undefined);
    message.submitTime !== undefined && (obj.submit_time = message.submitTime ?? undefined);
    message.depositEndTime !== undefined && (obj.deposit_end_time = message.depositEndTime ?? undefined);
    if (message.totalDeposit) {
      obj.total_deposit = message.totalDeposit.map(e => e ? Coin.toSDK(e) : undefined);
    } else {
      obj.total_deposit = [];
    }
    message.votingStartTime !== undefined && (obj.voting_start_time = message.votingStartTime ?? undefined);
    message.votingEndTime !== undefined && (obj.voting_end_time = message.votingEndTime ?? undefined);
    obj.metadata = message.metadata;
    return obj;
  },
  fromAmino(object: ProposalAmino): Proposal {
    return {
      id: Long.fromString(object.id),
      messages: Array.isArray(object?.messages) ? object.messages.map((e: any) => Any.fromAmino(e)) : [],
      status: isSet(object.status) ? proposalStatusFromJSON(object.status) : -1,
      finalTallyResult: object?.final_tally_result ? TallyResult.fromAmino(object.final_tally_result) : undefined,
      submitTime: object.submit_time,
      depositEndTime: object.deposit_end_time,
      totalDeposit: Array.isArray(object?.total_deposit) ? object.total_deposit.map((e: any) => Coin.fromAmino(e)) : [],
      votingStartTime: object.voting_start_time,
      votingEndTime: object.voting_end_time,
      metadata: object.metadata
    };
  },
  toAmino(message: Proposal): ProposalAmino {
    const obj: any = {};
    obj.id = message.id ? message.id.toString() : undefined;
    if (message.messages) {
      obj.messages = message.messages.map(e => e ? Any.toAmino(e) : undefined);
    } else {
      obj.messages = [];
    }
    obj.status = message.status;
    obj.final_tally_result = message.finalTallyResult ? TallyResult.toAmino(message.finalTallyResult) : undefined;
    obj.submit_time = message.submitTime;
    obj.deposit_end_time = message.depositEndTime;
    if (message.totalDeposit) {
      obj.total_deposit = message.totalDeposit.map(e => e ? Coin.toAmino(e) : undefined);
    } else {
      obj.total_deposit = [];
    }
    obj.voting_start_time = message.votingStartTime;
    obj.voting_end_time = message.votingEndTime;
    obj.metadata = message.metadata;
    return obj;
  },
  fromAminoMsg(object: ProposalAminoMsg): Proposal {
    return Proposal.fromAmino(object.value);
  },
  toAminoMsg(message: Proposal): ProposalAminoMsg {
    return {
      type: "cosmos-sdk/v1/Proposal",
      value: Proposal.toAmino(message)
    };
  },
  fromProtoMsg(message: ProposalProtoMsg): Proposal {
    return Proposal.decode(message.value);
  },
  toProto(message: Proposal): Uint8Array {
    return Proposal.encode(message).finish();
  },
  toProtoMsg(message: Proposal): ProposalProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.Proposal",
      value: Proposal.encode(message).finish()
    };
  }
};
function createBaseTallyResult(): TallyResult {
  return {
    yesCount: "",
    abstainCount: "",
    noCount: "",
    noWithVetoCount: ""
  };
}
export const TallyResult = {
  typeUrl: "/cosmos.gov.v1.TallyResult",
  aminoType: "cosmos-sdk/v1/TallyResult",
  encode(message: TallyResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.yesCount !== "") {
      writer.uint32(10).string(message.yesCount);
    }
    if (message.abstainCount !== "") {
      writer.uint32(18).string(message.abstainCount);
    }
    if (message.noCount !== "") {
      writer.uint32(26).string(message.noCount);
    }
    if (message.noWithVetoCount !== "") {
      writer.uint32(34).string(message.noWithVetoCount);
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): TallyResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTallyResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.yesCount = reader.string();
          break;
        case 2:
          message.abstainCount = reader.string();
          break;
        case 3:
          message.noCount = reader.string();
          break;
        case 4:
          message.noWithVetoCount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): TallyResult {
    const obj = createBaseTallyResult();
    if (isSet(object.yesCount)) obj.yesCount = String(object.yesCount);
    if (isSet(object.abstainCount)) obj.abstainCount = String(object.abstainCount);
    if (isSet(object.noCount)) obj.noCount = String(object.noCount);
    if (isSet(object.noWithVetoCount)) obj.noWithVetoCount = String(object.noWithVetoCount);
    return obj;
  },
  toJSON(message: TallyResult): unknown {
    const obj: any = {};
    message.yesCount !== undefined && (obj.yesCount = message.yesCount);
    message.abstainCount !== undefined && (obj.abstainCount = message.abstainCount);
    message.noCount !== undefined && (obj.noCount = message.noCount);
    message.noWithVetoCount !== undefined && (obj.noWithVetoCount = message.noWithVetoCount);
    return obj;
  },
  fromPartial(object: DeepPartial<TallyResult>): TallyResult {
    const message = createBaseTallyResult();
    message.yesCount = object.yesCount ?? "";
    message.abstainCount = object.abstainCount ?? "";
    message.noCount = object.noCount ?? "";
    message.noWithVetoCount = object.noWithVetoCount ?? "";
    return message;
  },
  fromSDK(object: TallyResultSDKType): TallyResult {
    return {
      yesCount: object?.yes_count,
      abstainCount: object?.abstain_count,
      noCount: object?.no_count,
      noWithVetoCount: object?.no_with_veto_count
    };
  },
  toSDK(message: TallyResult): TallyResultSDKType {
    const obj: any = {};
    obj.yes_count = message.yesCount;
    obj.abstain_count = message.abstainCount;
    obj.no_count = message.noCount;
    obj.no_with_veto_count = message.noWithVetoCount;
    return obj;
  },
  fromAmino(object: TallyResultAmino): TallyResult {
    return {
      yesCount: object.yes_count,
      abstainCount: object.abstain_count,
      noCount: object.no_count,
      noWithVetoCount: object.no_with_veto_count
    };
  },
  toAmino(message: TallyResult): TallyResultAmino {
    const obj: any = {};
    obj.yes_count = message.yesCount;
    obj.abstain_count = message.abstainCount;
    obj.no_count = message.noCount;
    obj.no_with_veto_count = message.noWithVetoCount;
    return obj;
  },
  fromAminoMsg(object: TallyResultAminoMsg): TallyResult {
    return TallyResult.fromAmino(object.value);
  },
  toAminoMsg(message: TallyResult): TallyResultAminoMsg {
    return {
      type: "cosmos-sdk/v1/TallyResult",
      value: TallyResult.toAmino(message)
    };
  },
  fromProtoMsg(message: TallyResultProtoMsg): TallyResult {
    return TallyResult.decode(message.value);
  },
  toProto(message: TallyResult): Uint8Array {
    return TallyResult.encode(message).finish();
  },
  toProtoMsg(message: TallyResult): TallyResultProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.TallyResult",
      value: TallyResult.encode(message).finish()
    };
  }
};
function createBaseVote(): Vote {
  return {
    proposalId: Long.UZERO,
    voter: "",
    options: [],
    metadata: ""
  };
}
export const Vote = {
  typeUrl: "/cosmos.gov.v1.Vote",
  aminoType: "cosmos-sdk/v1/Vote",
  encode(message: Vote, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.proposalId.isZero()) {
      writer.uint32(8).uint64(message.proposalId);
    }
    if (message.voter !== "") {
      writer.uint32(18).string(message.voter);
    }
    for (const v of message.options) {
      WeightedVoteOption.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.metadata !== "") {
      writer.uint32(42).string(message.metadata);
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): Vote {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVote();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposalId = (reader.uint64() as Long);
          break;
        case 2:
          message.voter = reader.string();
          break;
        case 4:
          message.options.push(WeightedVoteOption.decode(reader, reader.uint32()));
          break;
        case 5:
          message.metadata = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Vote {
    const obj = createBaseVote();
    if (isSet(object.proposalId)) obj.proposalId = Long.fromValue(object.proposalId);
    if (isSet(object.voter)) obj.voter = String(object.voter);
    if (Array.isArray(object?.options)) obj.options = object.options.map((e: any) => WeightedVoteOption.fromJSON(e));
    if (isSet(object.metadata)) obj.metadata = String(object.metadata);
    return obj;
  },
  toJSON(message: Vote): unknown {
    const obj: any = {};
    message.proposalId !== undefined && (obj.proposalId = (message.proposalId || Long.UZERO).toString());
    message.voter !== undefined && (obj.voter = message.voter);
    if (message.options) {
      obj.options = message.options.map(e => e ? WeightedVoteOption.toJSON(e) : undefined);
    } else {
      obj.options = [];
    }
    message.metadata !== undefined && (obj.metadata = message.metadata);
    return obj;
  },
  fromPartial(object: DeepPartial<Vote>): Vote {
    const message = createBaseVote();
    if (object.proposalId !== undefined && object.proposalId !== null) {
      message.proposalId = Long.fromValue(object.proposalId);
    }
    message.voter = object.voter ?? "";
    message.options = object.options?.map(e => WeightedVoteOption.fromPartial(e)) || [];
    message.metadata = object.metadata ?? "";
    return message;
  },
  fromSDK(object: VoteSDKType): Vote {
    return {
      proposalId: object?.proposal_id,
      voter: object?.voter,
      options: Array.isArray(object?.options) ? object.options.map((e: any) => WeightedVoteOption.fromSDK(e)) : [],
      metadata: object?.metadata
    };
  },
  toSDK(message: Vote): VoteSDKType {
    const obj: any = {};
    obj.proposal_id = message.proposalId;
    obj.voter = message.voter;
    if (message.options) {
      obj.options = message.options.map(e => e ? WeightedVoteOption.toSDK(e) : undefined);
    } else {
      obj.options = [];
    }
    obj.metadata = message.metadata;
    return obj;
  },
  fromAmino(object: VoteAmino): Vote {
    return {
      proposalId: Long.fromString(object.proposal_id),
      voter: object.voter,
      options: Array.isArray(object?.options) ? object.options.map((e: any) => WeightedVoteOption.fromAmino(e)) : [],
      metadata: object.metadata
    };
  },
  toAmino(message: Vote): VoteAmino {
    const obj: any = {};
    obj.proposal_id = message.proposalId ? message.proposalId.toString() : undefined;
    obj.voter = message.voter;
    if (message.options) {
      obj.options = message.options.map(e => e ? WeightedVoteOption.toAmino(e) : undefined);
    } else {
      obj.options = [];
    }
    obj.metadata = message.metadata;
    return obj;
  },
  fromAminoMsg(object: VoteAminoMsg): Vote {
    return Vote.fromAmino(object.value);
  },
  toAminoMsg(message: Vote): VoteAminoMsg {
    return {
      type: "cosmos-sdk/v1/Vote",
      value: Vote.toAmino(message)
    };
  },
  fromProtoMsg(message: VoteProtoMsg): Vote {
    return Vote.decode(message.value);
  },
  toProto(message: Vote): Uint8Array {
    return Vote.encode(message).finish();
  },
  toProtoMsg(message: Vote): VoteProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.Vote",
      value: Vote.encode(message).finish()
    };
  }
};
function createBaseDepositParams(): DepositParams {
  return {
    minDeposit: [],
    maxDepositPeriod: Duration.fromPartial({})
  };
}
export const DepositParams = {
  typeUrl: "/cosmos.gov.v1.DepositParams",
  aminoType: "cosmos-sdk/v1/DepositParams",
  encode(message: DepositParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.minDeposit) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.maxDepositPeriod !== undefined) {
      Duration.encode(message.maxDepositPeriod, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): DepositParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDepositParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minDeposit.push(Coin.decode(reader, reader.uint32()));
          break;
        case 2:
          message.maxDepositPeriod = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): DepositParams {
    const obj = createBaseDepositParams();
    if (Array.isArray(object?.minDeposit)) obj.minDeposit = object.minDeposit.map((e: any) => Coin.fromJSON(e));
    if (isSet(object.maxDepositPeriod)) obj.maxDepositPeriod = Duration.fromJSON(object.maxDepositPeriod);
    return obj;
  },
  toJSON(message: DepositParams): unknown {
    const obj: any = {};
    if (message.minDeposit) {
      obj.minDeposit = message.minDeposit.map(e => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.minDeposit = [];
    }
    message.maxDepositPeriod !== undefined && (obj.maxDepositPeriod = message.maxDepositPeriod ? Duration.toJSON(message.maxDepositPeriod) : undefined);
    return obj;
  },
  fromPartial(object: DeepPartial<DepositParams>): DepositParams {
    const message = createBaseDepositParams();
    message.minDeposit = object.minDeposit?.map(e => Coin.fromPartial(e)) || [];
    if (object.maxDepositPeriod !== undefined && object.maxDepositPeriod !== null) {
      message.maxDepositPeriod = Duration.fromPartial(object.maxDepositPeriod);
    }
    return message;
  },
  fromSDK(object: DepositParamsSDKType): DepositParams {
    return {
      minDeposit: Array.isArray(object?.min_deposit) ? object.min_deposit.map((e: any) => Coin.fromSDK(e)) : [],
      maxDepositPeriod: object.max_deposit_period ? Duration.fromSDK(object.max_deposit_period) : undefined
    };
  },
  toSDK(message: DepositParams): DepositParamsSDKType {
    const obj: any = {};
    if (message.minDeposit) {
      obj.min_deposit = message.minDeposit.map(e => e ? Coin.toSDK(e) : undefined);
    } else {
      obj.min_deposit = [];
    }
    message.maxDepositPeriod !== undefined && (obj.max_deposit_period = message.maxDepositPeriod ? Duration.toSDK(message.maxDepositPeriod) : undefined);
    return obj;
  },
  fromAmino(object: DepositParamsAmino): DepositParams {
    return {
      minDeposit: Array.isArray(object?.min_deposit) ? object.min_deposit.map((e: any) => Coin.fromAmino(e)) : [],
      maxDepositPeriod: object?.max_deposit_period ? Duration.fromAmino(object.max_deposit_period) : undefined
    };
  },
  toAmino(message: DepositParams): DepositParamsAmino {
    const obj: any = {};
    if (message.minDeposit) {
      obj.min_deposit = message.minDeposit.map(e => e ? Coin.toAmino(e) : undefined);
    } else {
      obj.min_deposit = [];
    }
    obj.max_deposit_period = message.maxDepositPeriod ? Duration.toAmino(message.maxDepositPeriod) : undefined;
    return obj;
  },
  fromAminoMsg(object: DepositParamsAminoMsg): DepositParams {
    return DepositParams.fromAmino(object.value);
  },
  toAminoMsg(message: DepositParams): DepositParamsAminoMsg {
    return {
      type: "cosmos-sdk/v1/DepositParams",
      value: DepositParams.toAmino(message)
    };
  },
  fromProtoMsg(message: DepositParamsProtoMsg): DepositParams {
    return DepositParams.decode(message.value);
  },
  toProto(message: DepositParams): Uint8Array {
    return DepositParams.encode(message).finish();
  },
  toProtoMsg(message: DepositParams): DepositParamsProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.DepositParams",
      value: DepositParams.encode(message).finish()
    };
  }
};
function createBaseVotingParams(): VotingParams {
  return {
    votingPeriod: Duration.fromPartial({})
  };
}
export const VotingParams = {
  typeUrl: "/cosmos.gov.v1.VotingParams",
  aminoType: "cosmos-sdk/v1/VotingParams",
  encode(message: VotingParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.votingPeriod !== undefined) {
      Duration.encode(message.votingPeriod, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): VotingParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVotingParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.votingPeriod = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): VotingParams {
    const obj = createBaseVotingParams();
    if (isSet(object.votingPeriod)) obj.votingPeriod = Duration.fromJSON(object.votingPeriod);
    return obj;
  },
  toJSON(message: VotingParams): unknown {
    const obj: any = {};
    message.votingPeriod !== undefined && (obj.votingPeriod = message.votingPeriod ? Duration.toJSON(message.votingPeriod) : undefined);
    return obj;
  },
  fromPartial(object: DeepPartial<VotingParams>): VotingParams {
    const message = createBaseVotingParams();
    if (object.votingPeriod !== undefined && object.votingPeriod !== null) {
      message.votingPeriod = Duration.fromPartial(object.votingPeriod);
    }
    return message;
  },
  fromSDK(object: VotingParamsSDKType): VotingParams {
    return {
      votingPeriod: object.voting_period ? Duration.fromSDK(object.voting_period) : undefined
    };
  },
  toSDK(message: VotingParams): VotingParamsSDKType {
    const obj: any = {};
    message.votingPeriod !== undefined && (obj.voting_period = message.votingPeriod ? Duration.toSDK(message.votingPeriod) : undefined);
    return obj;
  },
  fromAmino(object: VotingParamsAmino): VotingParams {
    return {
      votingPeriod: object?.voting_period ? Duration.fromAmino(object.voting_period) : undefined
    };
  },
  toAmino(message: VotingParams): VotingParamsAmino {
    const obj: any = {};
    obj.voting_period = message.votingPeriod ? Duration.toAmino(message.votingPeriod) : undefined;
    return obj;
  },
  fromAminoMsg(object: VotingParamsAminoMsg): VotingParams {
    return VotingParams.fromAmino(object.value);
  },
  toAminoMsg(message: VotingParams): VotingParamsAminoMsg {
    return {
      type: "cosmos-sdk/v1/VotingParams",
      value: VotingParams.toAmino(message)
    };
  },
  fromProtoMsg(message: VotingParamsProtoMsg): VotingParams {
    return VotingParams.decode(message.value);
  },
  toProto(message: VotingParams): Uint8Array {
    return VotingParams.encode(message).finish();
  },
  toProtoMsg(message: VotingParams): VotingParamsProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.VotingParams",
      value: VotingParams.encode(message).finish()
    };
  }
};
function createBaseTallyParams(): TallyParams {
  return {
    quorum: "",
    threshold: "",
    vetoThreshold: ""
  };
}
export const TallyParams = {
  typeUrl: "/cosmos.gov.v1.TallyParams",
  aminoType: "cosmos-sdk/v1/TallyParams",
  encode(message: TallyParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.quorum !== "") {
      writer.uint32(10).string(message.quorum);
    }
    if (message.threshold !== "") {
      writer.uint32(18).string(message.threshold);
    }
    if (message.vetoThreshold !== "") {
      writer.uint32(26).string(message.vetoThreshold);
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): TallyParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTallyParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.quorum = reader.string();
          break;
        case 2:
          message.threshold = reader.string();
          break;
        case 3:
          message.vetoThreshold = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): TallyParams {
    const obj = createBaseTallyParams();
    if (isSet(object.quorum)) obj.quorum = String(object.quorum);
    if (isSet(object.threshold)) obj.threshold = String(object.threshold);
    if (isSet(object.vetoThreshold)) obj.vetoThreshold = String(object.vetoThreshold);
    return obj;
  },
  toJSON(message: TallyParams): unknown {
    const obj: any = {};
    message.quorum !== undefined && (obj.quorum = message.quorum);
    message.threshold !== undefined && (obj.threshold = message.threshold);
    message.vetoThreshold !== undefined && (obj.vetoThreshold = message.vetoThreshold);
    return obj;
  },
  fromPartial(object: DeepPartial<TallyParams>): TallyParams {
    const message = createBaseTallyParams();
    message.quorum = object.quorum ?? "";
    message.threshold = object.threshold ?? "";
    message.vetoThreshold = object.vetoThreshold ?? "";
    return message;
  },
  fromSDK(object: TallyParamsSDKType): TallyParams {
    return {
      quorum: object?.quorum,
      threshold: object?.threshold,
      vetoThreshold: object?.veto_threshold
    };
  },
  toSDK(message: TallyParams): TallyParamsSDKType {
    const obj: any = {};
    obj.quorum = message.quorum;
    obj.threshold = message.threshold;
    obj.veto_threshold = message.vetoThreshold;
    return obj;
  },
  fromAmino(object: TallyParamsAmino): TallyParams {
    return {
      quorum: object.quorum,
      threshold: object.threshold,
      vetoThreshold: object.veto_threshold
    };
  },
  toAmino(message: TallyParams): TallyParamsAmino {
    const obj: any = {};
    obj.quorum = message.quorum;
    obj.threshold = message.threshold;
    obj.veto_threshold = message.vetoThreshold;
    return obj;
  },
  fromAminoMsg(object: TallyParamsAminoMsg): TallyParams {
    return TallyParams.fromAmino(object.value);
  },
  toAminoMsg(message: TallyParams): TallyParamsAminoMsg {
    return {
      type: "cosmos-sdk/v1/TallyParams",
      value: TallyParams.toAmino(message)
    };
  },
  fromProtoMsg(message: TallyParamsProtoMsg): TallyParams {
    return TallyParams.decode(message.value);
  },
  toProto(message: TallyParams): Uint8Array {
    return TallyParams.encode(message).finish();
  },
  toProtoMsg(message: TallyParams): TallyParamsProtoMsg {
    return {
      typeUrl: "/cosmos.gov.v1.TallyParams",
      value: TallyParams.encode(message).finish()
    };
  }
};