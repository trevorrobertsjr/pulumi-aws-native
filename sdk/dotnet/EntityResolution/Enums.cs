// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.EntityResolution
{
    [EnumType]
    public readonly struct IdMappingWorkflowIdMappingTechniquesIdMappingType : IEquatable<IdMappingWorkflowIdMappingTechniquesIdMappingType>
    {
        private readonly string _value;

        private IdMappingWorkflowIdMappingTechniquesIdMappingType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdMappingWorkflowIdMappingTechniquesIdMappingType Provider { get; } = new IdMappingWorkflowIdMappingTechniquesIdMappingType("PROVIDER");

        public static bool operator ==(IdMappingWorkflowIdMappingTechniquesIdMappingType left, IdMappingWorkflowIdMappingTechniquesIdMappingType right) => left.Equals(right);
        public static bool operator !=(IdMappingWorkflowIdMappingTechniquesIdMappingType left, IdMappingWorkflowIdMappingTechniquesIdMappingType right) => !left.Equals(right);

        public static explicit operator string(IdMappingWorkflowIdMappingTechniquesIdMappingType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdMappingWorkflowIdMappingTechniquesIdMappingType other && Equals(other);
        public bool Equals(IdMappingWorkflowIdMappingTechniquesIdMappingType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct MatchingWorkflowResolutionTechniquesResolutionType : IEquatable<MatchingWorkflowResolutionTechniquesResolutionType>
    {
        private readonly string _value;

        private MatchingWorkflowResolutionTechniquesResolutionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MatchingWorkflowResolutionTechniquesResolutionType RuleMatching { get; } = new MatchingWorkflowResolutionTechniquesResolutionType("RULE_MATCHING");
        public static MatchingWorkflowResolutionTechniquesResolutionType MlMatching { get; } = new MatchingWorkflowResolutionTechniquesResolutionType("ML_MATCHING");
        public static MatchingWorkflowResolutionTechniquesResolutionType Provider { get; } = new MatchingWorkflowResolutionTechniquesResolutionType("PROVIDER");

        public static bool operator ==(MatchingWorkflowResolutionTechniquesResolutionType left, MatchingWorkflowResolutionTechniquesResolutionType right) => left.Equals(right);
        public static bool operator !=(MatchingWorkflowResolutionTechniquesResolutionType left, MatchingWorkflowResolutionTechniquesResolutionType right) => !left.Equals(right);

        public static explicit operator string(MatchingWorkflowResolutionTechniquesResolutionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MatchingWorkflowResolutionTechniquesResolutionType other && Equals(other);
        public bool Equals(MatchingWorkflowResolutionTechniquesResolutionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel : IEquatable<MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel>
    {
        private readonly string _value;

        private MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel OneToOne { get; } = new MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel("ONE_TO_ONE");
        public static MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel ManyToMany { get; } = new MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel("MANY_TO_MANY");

        public static bool operator ==(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel left, MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel right) => left.Equals(right);
        public static bool operator !=(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel left, MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel right) => !left.Equals(right);

        public static explicit operator string(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel other && Equals(other);
        public bool Equals(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SchemaMappingSchemaAttributeType : IEquatable<SchemaMappingSchemaAttributeType>
    {
        private readonly string _value;

        private SchemaMappingSchemaAttributeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SchemaMappingSchemaAttributeType Name { get; } = new SchemaMappingSchemaAttributeType("NAME");
        public static SchemaMappingSchemaAttributeType NameFirst { get; } = new SchemaMappingSchemaAttributeType("NAME_FIRST");
        public static SchemaMappingSchemaAttributeType NameMiddle { get; } = new SchemaMappingSchemaAttributeType("NAME_MIDDLE");
        public static SchemaMappingSchemaAttributeType NameLast { get; } = new SchemaMappingSchemaAttributeType("NAME_LAST");
        public static SchemaMappingSchemaAttributeType Address { get; } = new SchemaMappingSchemaAttributeType("ADDRESS");
        public static SchemaMappingSchemaAttributeType AddressStreet1 { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_STREET1");
        public static SchemaMappingSchemaAttributeType AddressStreet2 { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_STREET2");
        public static SchemaMappingSchemaAttributeType AddressStreet3 { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_STREET3");
        public static SchemaMappingSchemaAttributeType AddressCity { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_CITY");
        public static SchemaMappingSchemaAttributeType AddressState { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_STATE");
        public static SchemaMappingSchemaAttributeType AddressCountry { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_COUNTRY");
        public static SchemaMappingSchemaAttributeType AddressPostalcode { get; } = new SchemaMappingSchemaAttributeType("ADDRESS_POSTALCODE");
        public static SchemaMappingSchemaAttributeType Phone { get; } = new SchemaMappingSchemaAttributeType("PHONE");
        public static SchemaMappingSchemaAttributeType PhoneNumber { get; } = new SchemaMappingSchemaAttributeType("PHONE_NUMBER");
        public static SchemaMappingSchemaAttributeType PhoneCountrycode { get; } = new SchemaMappingSchemaAttributeType("PHONE_COUNTRYCODE");
        public static SchemaMappingSchemaAttributeType EmailAddress { get; } = new SchemaMappingSchemaAttributeType("EMAIL_ADDRESS");
        public static SchemaMappingSchemaAttributeType UniqueId { get; } = new SchemaMappingSchemaAttributeType("UNIQUE_ID");
        public static SchemaMappingSchemaAttributeType Date { get; } = new SchemaMappingSchemaAttributeType("DATE");
        public static SchemaMappingSchemaAttributeType String { get; } = new SchemaMappingSchemaAttributeType("STRING");
        public static SchemaMappingSchemaAttributeType ProviderId { get; } = new SchemaMappingSchemaAttributeType("PROVIDER_ID");

        public static bool operator ==(SchemaMappingSchemaAttributeType left, SchemaMappingSchemaAttributeType right) => left.Equals(right);
        public static bool operator !=(SchemaMappingSchemaAttributeType left, SchemaMappingSchemaAttributeType right) => !left.Equals(right);

        public static explicit operator string(SchemaMappingSchemaAttributeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SchemaMappingSchemaAttributeType other && Equals(other);
        public bool Equals(SchemaMappingSchemaAttributeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
